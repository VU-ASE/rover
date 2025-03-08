package commands

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/signal"
	"strings"
	"syscall"

	proxy "github.com/VU-ASE/rover/roverctl/src/proxy"
	style "github.com/VU-ASE/rover/roverctl/src/style"
	"github.com/VU-ASE/rover/roverctl/src/utils"

	command_prechecks "github.com/VU-ASE/rover/roverctl/src/commands/prechecks"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/api/types/registry"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"github.com/spf13/cobra"
)

func NewRoot() *cobra.Command {
	// General flags
	var roverIndex int
	var roverdHost string
	var roverdUsername string
	var roverdPassword string

	//
	// CLI commands
	//
	var debugMode bool         // enable proxy server and inject its IP into roverctl-web
	var verbose bool           // show debug logs
	var roverctlVersion string // force roverctl at a specific version
	var rootCmd = &cobra.Command{
		Use:   "roverctl",
		Short: "CLI to manage a Rover",
		Long:  "A command line interface to manage a Rover",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			// Check if Docker is running
			cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
			if err != nil {
				return fmt.Errorf("failed to initialize Docker client: %v", err)
			}
			_, err = cli.Ping(context.Background())
			if err != nil {
				return fmt.Errorf("Docker daemon not reachable: %v", err)
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			conn, err := command_prechecks.Perform(cmd, args, roverIndex, roverdHost, roverdUsername, roverdPassword)
			if err != nil {
				return err
			}

			// If version is set, make sure it has the "v" prefix
			if roverctlVersion != "" {
				roverctlVersion = "v" + strings.TrimPrefix(roverctlVersion, "v")
			}

			//
			// Check what version roverd is running on the Rover
			//

			version := ""
			if roverctlVersion == "" {
				fmt.Print("Connecting to Rover to determine roverctl-web version to use...\n")
				api := conn.ToApiClient()
				res, _, err := api.HealthAPI.StatusGet(
					context.Background(),
				).Execute()
				if err != nil {
					fmt.Printf("Failed to connect to Rover: %v", err)
					return nil
				}
				version = "v" + strings.TrimPrefix(res.Version, "v")
				fmt.Printf("Rover is running roverd version %s\n", style.Success.Render(version))
			} else {
				fmt.Printf("Forcing roverctl-web to run at version %s\n", style.Success.Render(roverctlVersion))
				version = roverctlVersion
			}

			//
			// Find out if there is a matching roverctl-web version
			//

			// Initialize Docker client
			dc, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
			if err != nil {
				return fmt.Errorf("failed to create Docker client: %v", err)
			}
			ctx := context.Background()
			dc.NegotiateAPIVersion(ctx)

			// Register with GHCR
			author := "vu-ase"
			name := "roverctl-web"
			imageRef := fmt.Sprintf("ghcr.io/%s/%s:%s", author, name, version)
			ghcr := registry.AuthConfig{
				ServerAddress: "ghcr.io",
			}
			encodedAuth, err := registry.EncodeAuthConfig(ghcr)
			if err != nil {
				return fmt.Errorf("failed to encode auth: %v", err)
			}

			// Check if image exists
			_, err = dc.DistributionInspect(ctx, imageRef, encodedAuth)
			if err != nil {
				fmt.Printf("No matching roverctl-web image found for roverd version %s. Upgrade roverd to ensure compatibility or use the --force flag to use a specific version.\n", version)
				return nil
			}

			// Pull roverctl-web
			fmt.Printf("Found matching roverctl-web image, pulling...\n")
			out, err := dc.ImagePull(ctx, imageRef, image.PullOptions{})
			if err != nil {
				fmt.Println("Error pulling image:", err)
				return nil
			}
			defer out.Close()
			if verbose {
				io.Copy(os.Stdout, out) // Stream pull output to console
			}

			//
			// Start the container(s)
			//

			proxyHost := ""
			if debugMode {
				proxyHost, err = utils.GetLocalIP()
				if err != nil {
					fmt.Println("Failed to get local IP necessary for debugging:", err)
					return nil
				}
				fmt.Printf("Using local IP %s for debugging\n", proxyHost)
			}

			proxyHttpPort := 7500
			proxyUdpPort := 40000
			// Environment variables roverctl-web needs
			envVars := []string{
				"PUBLIC_ROVERD_HOST=" + conn.Host,
				"PUBLIC_ROVERD_PORT=80",
				"PUBLIC_ROVERD_USERNAME=" + conn.Username,
				"PUBLIC_ROVERD_PASSWORD=" + conn.Password,
				"PUBLIC_PASSTHROUGH_HOST=" + proxyHost,
				"PUBLIC_PASSTHROUGH_PORT=" + fmt.Sprintf("%d", proxyHttpPort),
			}

			// Port forwarding (host:container)
			portBindings := nat.PortMap{
				"3000/tcp": []nat.PortBinding{
					{HostIP: "0.0.0.0", HostPort: "3000"},
				},
			}

			// Create a container with the specified image and environment variables
			resp, err := dc.ContainerCreate(ctx, &container.Config{
				Image: imageRef,
				Env:   envVars,
				Tty:   true, // Keep terminal session interactive
				ExposedPorts: nat.PortSet{
					"3000/tcp": struct{}{}, // Expose container's port 80
				},
			}, &container.HostConfig{
				PortBindings: portBindings, // Port forwarding
			}, &network.NetworkingConfig{}, nil, "")
			if err != nil {
				fmt.Printf("failed to create roverctl-web container: %v", err)
				return nil
			}

			// Run the proxy server
			if debugMode {
				go proxy.Run(proxyHost, proxyUdpPort, false, verbose)
			}

			// Start the container
			if err := dc.ContainerStart(ctx, resp.ID, container.StartOptions{}); err != nil {
				fmt.Printf("failed to start roverctl-web container: %v", err)
				return nil
			}

			url := "http://localhost:3000"
			fmt.Printf("Visit %s to control this Rover!\n", style.Primary.Render(url))
			utils.OpenBrowser(url)

			// Set up signal handling (to stop container on Ctrl+C)
			sigChan := make(chan os.Signal, 1)
			signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

			go func() {
				<-sigChan // Wait for Ctrl+C (SIGINT)
				fmt.Println("\nStopping roverctl-web...")

				// Stop container with a timeout (graceful shutdown)
				timeout := 10 // seconds
				err := dc.ContainerStop(ctx, resp.ID, container.StopOptions{Timeout: &timeout})
				if err != nil {
					fmt.Printf("failed to stop roverctl-web container: %v\n", err)
				} else {
					fmt.Println("roverctl-web container stopped successfully")
				}

				os.Exit(0)
			}()

			// Attach to container (to keep it running)
			statusCh, errCh := dc.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
			select {
			case err := <-errCh:
				if err != nil {
					fmt.Printf("roverctl-web container error: %v\n", err)
				}
			case <-statusCh:
			}

			fmt.Println("roverctl-web container exited")
			return nil
		},
	}
	rootCmd.Flags().IntVarP(&roverIndex, "rover", "r", 0, "The index of the rover to upload to, between 1 and 20")
	rootCmd.Flags().StringVarP(&roverdHost, "host", "", "", "The roverd endpoint to connect to (if not using --rover)")
	rootCmd.Flags().StringVarP(&roverdUsername, "username", "u", "debix", "The username to use to connect to the roverd endpoint")
	rootCmd.Flags().StringVarP(&roverdPassword, "password", "p", "debix", "The password to use to connect to the roverd endpoint")
	rootCmd.Flags().BoolVarP(&debugMode, "debug", "d", false, "Enable debug/tuning mode")
	rootCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Show verbose output")
	rootCmd.Flags().StringVarP(&roverctlVersion, "force", "f", "", "Force roverctl-web to run at a specific version")

	return rootCmd
}
