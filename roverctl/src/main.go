package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"

	"github.com/VU-ASE/rover/roverctl/src/configuration"
	proxy "github.com/VU-ASE/rover/roverctl/src/proxy"
	"github.com/VU-ASE/rover/roverctl/src/state"
	style "github.com/VU-ASE/rover/roverctl/src/style"
	"github.com/VU-ASE/rover/roverctl/src/utils"
	view_info "github.com/VU-ASE/rover/roverctl/src/views/info"

	view_upload "github.com/VU-ASE/rover/roverctl/src/views/upload"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/api/types/registry"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func prechecks(cmd *cobra.Command, args []string, roverIndex int, roverdHost string, roverdUsername, roverdPassword string) (*configuration.RoverConnection, error) {
	// XOR Logic: Exactly one must be set
	roverSet := cmd.Flags().Changed("rover")
	hostSet := cmd.Flags().Changed("host")
	if roverSet == hostSet { // both false or both true
		return nil, fmt.Errorf("you must provide either --rover or --host, but not both")
	}

	identifier := roverdHost
	host := roverdHost
	if roverSet {
		if roverIndex < 1 || roverIndex > 20 {
			return nil, fmt.Errorf("rover index must be between 1 and 20")
		}
		identifier = fmt.Sprintf("rover %d", roverIndex)
		host = fmt.Sprintf("192.168.0.%d", roverIndex+100)
	}

	// Create connection
	conn := configuration.RoverConnection{
		Identifier: identifier,
		Host:       host,
		Username:   roverdUsername,
		Password:   roverdPassword,
	}
	return &conn, nil
}

func run() error {
	// Initialize the app and create app state
	err := configuration.Initialize()
	if err != nil {
		return err
	}

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
			conn, err := prechecks(cmd, args, roverIndex, roverdHost, roverdUsername, roverdPassword)
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

	// Upload command
	var watch bool
	var uploadCmd = &cobra.Command{
		Use:   "upload <PATHS>",
		Short: "Upload specified service folders to a Rover",
		Long: `The upload command allows you to upload one or more service folders to the Rover. 
You can optionally specify the --watch flag to enable file watch and upload.`,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return fmt.Errorf("at least one directory must be provided")
			}
			for _, dir := range args {
				info, err := os.Stat(dir)
				if err != nil {
					return fmt.Errorf("invalid directory '%s': %v", dir, err)
				}
				if !info.IsDir() {
					return fmt.Errorf("'%s' is not a directory", dir)
				}
			}
			return nil
		},

		RunE: func(cmd *cobra.Command, args []string) error {
			conn, err := prechecks(cmd, args, roverIndex, roverdHost, roverdUsername, roverdPassword)
			if err != nil {
				return err
			}

			p := tea.NewProgram(view_upload.New(
				*conn, args, watch,
			))
			_, err = p.Run()
			return err
		},
	}
	uploadCmd.Flags().BoolVarP(&watch, "watch", "w", false, "Enable file watching")
	uploadCmd.Flags().IntVarP(&roverIndex, "rover", "r", 0, "The index of the rover to upload to, between 1 and 20")
	uploadCmd.Flags().StringVarP(&roverdHost, "host", "", "", "The roverd endpoint to connect to (if not using --rover)")
	uploadCmd.Flags().StringVarP(&roverdUsername, "username", "u", "debix", "The username to use to connect to the roverd endpoint")
	uploadCmd.Flags().StringVarP(&roverdPassword, "password", "p", "debix", "The password to use to connect to the roverd endpoint")
	rootCmd.AddCommand(uploadCmd)

	// SSH command
	var sshCmd = &cobra.Command{
		Use:   "ssh",
		Short: "Open an SSH terminal to a Rover",
		Long:  `Will use native SSH to open a terminal to the Rover.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			conn, err := prechecks(cmd, args, roverIndex, roverdHost, roverdUsername, roverdPassword)
			if err != nil {
				return err
			}

			sshPath, err := exec.LookPath("ssh") // Find the SSH binary
			if err != nil {
				fmt.Println("Could not find ssh in PATH")
			}

			host := fmt.Sprintf("%s@%s", conn.Username, conn.Host)
			err = syscall.Exec(sshPath, []string{"ssh", host}, os.Environ())
			if err != nil {
				return err
			}
			return nil
		},
	}
	sshCmd.Flags().IntVarP(&roverIndex, "rover", "r", 0, "The index of the rover to upload to, between 1 and 20")
	sshCmd.Flags().StringVarP(&roverdHost, "host", "", "", "The roverd endpoint to connect to (if not using --rover)")
	sshCmd.Flags().StringVarP(&roverdUsername, "username", "u", "debix", "The username to use to connect to the roverd endpoint")
	sshCmd.Flags().StringVarP(&roverdPassword, "password", "p", "debix", "The password to use to connect to the roverd endpoint")
	rootCmd.AddCommand(sshCmd)

	// info command
	var infoCmd = &cobra.Command{
		Use:   "info",
		Short: "View roverctl and roverd information",
		Long:  `Display build and connection information for roverctl, and roverd if a rover is specified.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			// Ignore errors
			conn, _ := prechecks(cmd, args, roverIndex, roverdHost, roverdUsername, roverdPassword)

			p := tea.NewProgram(view_info.New(conn))
			_, err = p.Run()
			return err
		},
	}
	infoCmd.Flags().IntVarP(&roverIndex, "rover", "r", 0, "The index of the rover to upload to, between 1 and 20")
	infoCmd.Flags().StringVarP(&roverdHost, "host", "", "", "The roverd endpoint to connect to (if not using --rover)")
	infoCmd.Flags().StringVarP(&roverdUsername, "username", "u", "debix", "The username to use to connect to the roverd endpoint")
	infoCmd.Flags().StringVarP(&roverdPassword, "password", "p", "debix", "The password to use to connect to the roverd endpoint")
	rootCmd.AddCommand(infoCmd)

	// Self-update command
	var version string
	var selfUpdateCmd = &cobra.Command{
		Use:   "update",
		Short: "Self-update roverctl",
		Long:  `Update roverctl to the latest version, or the version specified.`,
		Run: func(cmd *cobra.Command, args []string) {
			if version != "" {
				version = strings.TrimPrefix(version, "v")
				utils.ExecuteShellCommand("curl -fsSL https://raw.githubusercontent.com/VU-ASE/rover/refs/heads/main/roverctl/install.sh | bash -s v" + version)
			} else {
				utils.ExecuteShellCommand("curl -fsSL https://raw.githubusercontent.com/VU-ASE/rover/refs/heads/main/roverctl/install.sh | bash")
			}
		},
	}
	selfUpdateCmd.Flags().StringVarP(&version, "version", "v", "", "The version tag to update/downgrade to (e.g. v0.1.0)")
	rootCmd.AddCommand(selfUpdateCmd)

	// state.Get().QuitCommand =

	err = rootCmd.Execute()
	if err != nil {
		log.Err(err)
		// don't return, clean up
	}

	// Save configs to disk
	err = state.Get().Config.Save()
	if err != nil {
		return err
	}

	quitCmd := state.Get().QuitCommand
	if quitCmd != "" {
		return utils.ExecuteShellCommand(quitCmd)
	}

	return nil
}

func main() {
	// Configure zerolog to output to stdout beautifully
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}) //

	// Run the program
	if err := run(); err != nil {
		log.Error().Err(err).Msg("An error occurred while running the program.")
		os.Exit(1)
	}
}
