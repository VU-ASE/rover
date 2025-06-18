package commands

import (
	"context"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/VU-ASE/rover/roverctl/src/configuration"
	proxy "github.com/VU-ASE/rover/roverctl/src/proxy"
	style "github.com/VU-ASE/rover/roverctl/src/style"
	"github.com/VU-ASE/rover/roverctl/src/utils"
	view_info "github.com/VU-ASE/rover/roverctl/src/views/info"

	command_prechecks "github.com/VU-ASE/rover/roverctl/src/commands/prechecks"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/api/types/registry"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"github.com/spf13/cobra"
)

// checkInternetConnectivity checks if the system has internet access
func checkInternetConnectivity() bool {
	timeout := time.Duration(5 * time.Second)
	conn, err := net.DialTimeout("tcp", "8.8.8.8:53", timeout)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

// listCachedRoverctlImages returns a list of cached roverctl-web images
func listCachedRoverctlImages(dc *client.Client, ctx context.Context) ([]string, error) {
	images, err := dc.ImageList(ctx, image.ListOptions{})
	if err != nil {
		return nil, err
	}

	var roverctlImages []string
	for _, img := range images {
		for _, tag := range img.RepoTags {
			if strings.Contains(tag, "roverctl-web") {
				roverctlImages = append(roverctlImages, tag)
			}
		}
	}
	return roverctlImages, nil
}

// imageExists checks if a specific Docker image exists locally
func imageExists(dc *client.Client, ctx context.Context, imageRef string) bool {
	_, _, err := dc.ImageInspectWithRaw(ctx, imageRef)
	return err == nil
}

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
	var proxyIp string         // if a custom IP is specified, use that instead of the detected local IP
	var rootCmd = &cobra.Command{
		Use:   "roverctl",
		Short: "CLI to manage a Rover",
		Long:  "A command line interface to manage a Rover",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			// Check if Docker is running - try multiple socket locations for macOS compatibility
			var cli *client.Client
			var err error

			// First try the default (uses DOCKER_HOST env var or default socket)
			cli, err = client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
			if err == nil {
				_, err = cli.Ping(context.Background())
				if err == nil {
					return nil // Successfully connected
				}
			}

			// If default failed, try common macOS Docker Desktop socket locations
			possibleSockets := []string{
				"unix:///var/run/docker.sock",                                     // Standard Linux location
				"unix://" + os.Getenv("HOME") + "/.docker/run/docker.sock",        // macOS Docker Desktop
				"unix:///Users/" + os.Getenv("USER") + "/.docker/run/docker.sock", // Alternative macOS location
			}

			for _, socket := range possibleSockets {
				cli, err = client.NewClientWithOpts(client.WithHost(socket), client.WithAPIVersionNegotiation())
				if err == nil {
					_, err = cli.Ping(context.Background())
					if err == nil {
						// Set DOCKER_HOST for subsequent calls
						os.Setenv("DOCKER_HOST", socket)
						return nil
					}
				}
			}

			return fmt.Errorf("Docker daemon not reachable at any known location. Tried: %v. Last error: %v", possibleSockets, err)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			// Initialize Docker client
			dc, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
			if err != nil {
				return fmt.Errorf("failed to create Docker client: %v", err)
			}
			ctx := context.Background()
			dc.NegotiateAPIVersion(ctx)

			// If version is set, make sure it has the "v" prefix
			if roverctlVersion != "" {
				roverctlVersion = "v" + strings.TrimPrefix(roverctlVersion, "v")
			}

			var version string
			var conn *configuration.RoverConnection

			//
			// Determine version to use
			//
			if roverctlVersion == "" {
				// Connect to rover to get version (this works on LAN regardless of internet)
				conn, err = command_prechecks.Perform(cmd, args, roverIndex, roverdHost, roverdUsername, roverdPassword)
				if err != nil {
					return err
				}

				fmt.Print("Connecting to Rover to determine roverctl-web version to use...\n")

				api := conn.ToApiClient()
				res, _, err := api.HealthAPI.StatusGet(
					context.Background(),
				).Execute()
				if err != nil {
					fmt.Printf("Failed to connect to Rover: %v", err)
					return nil
				}
				version = utils.Version(res.Version)

				if !utils.VersionsEqual(version, view_info.Version) {
					fmt.Printf("Rover is running roverd %s but you are running roverctl %s\n", style.Warning.Render(version), style.Warning.Render(utils.Version(view_info.Version)))

					flagSuffix := ""
					if roverIndex > 0 {
						flagSuffix = fmt.Sprintf("--rover %d", roverIndex)
					} else {
						flagSuffix = fmt.Sprintf("--host %s", roverdHost)
					}

					if roverdUsername != "debix" {
						flagSuffix += fmt.Sprintf(" --username %s", roverdUsername)
					}
					if roverdPassword != "debix" {
						flagSuffix += fmt.Sprintf(" --password %s", roverdPassword)
					}

					fmt.Printf("\nThe following fixes are available:\n")
					fmt.Printf(" %s\n", style.Primary.Render("Update both roverctl and roverd to the latest version (recommended)"))
					fmt.Printf("   %s\n", style.Gray.Render("roverctl update "+flagSuffix))
					fmt.Printf(" %s\n", style.Primary.Render("OR try to match roverctl to the roverd version"))
					fmt.Printf("   %s\n", style.Gray.Render("roverctl update roverctl -v "+version))
					fmt.Printf(" %s\n", style.Primary.Render("OR try to match roverd to the roverctl version"))
					fmt.Printf("   %s\n", style.Gray.Render("roverctl update roverd -v "+utils.Version(view_info.Version)+" "+flagSuffix))
					fmt.Printf(" %s\n", style.Primary.Render("OR force roverctl to run at the roverd version"))
					fmt.Printf("   %s\n", style.Gray.Render("roverctl --force "+version+" "+flagSuffix))

					return nil
				} else {
					fmt.Printf("Rover is running roverd %s\n", style.Success.Render(version))
				}
			} else {
				fmt.Printf("Forcing roverctl-web to run at version %s\n", style.Success.Render(roverctlVersion))
				version = roverctlVersion
				// Still need connection details for container environment variables
				conn, err = command_prechecks.Perform(cmd, args, roverIndex, roverdHost, roverdUsername, roverdPassword)
				if err != nil {
					return err
				}
			}

			//
			// Check if image exists locally or pull it
			//
			author := "vu-ase"
			name := "roverctl-web"
			imageRef := fmt.Sprintf("ghcr.io/%s/%s:%s", author, name, version)

			// Check if image exists locally first
			if imageExists(dc, ctx, imageRef) {
				// Image exists locally - check if we're offline to show appropriate message
				hasInternet := checkInternetConnectivity()
				if !hasInternet {
					fmt.Printf("%s Using cached roverctl-web %s image (no internet connection)\n",
						style.Warning.Render("⚠"), style.Success.Render(version))
				} else {
					fmt.Printf("Using cached roverctl-web %s image\n", style.Success.Render(version))
				}
			} else {
				// Image doesn't exist locally - check internet connectivity for pulling
				hasInternet := checkInternetConnectivity()
				if !hasInternet {
					fmt.Printf("%s roverctl-web %s image not found locally and no internet connection available.\n",
						style.Warning.Render("⚠"), version)

					cachedImages, err := listCachedRoverctlImages(dc, ctx)
					if err != nil {
						fmt.Printf("Failed to list cached images: %v\n", err)
						return nil
					}

					if len(cachedImages) > 0 {
						fmt.Printf("\nAvailable cached versions:\n")
						for _, img := range cachedImages {
							parts := strings.Split(img, ":")
							if len(parts) > 1 {
								imgVersion := parts[len(parts)-1]
								fmt.Printf("  %s\n", style.Primary.Render(imgVersion))
							}
						}
						fmt.Printf("\nUse --force <version> to run with a cached version.\n")
						// Extract version from first cached image for example
						if len(cachedImages) > 0 {
							exampleVersion := cachedImages[0]
							if idx := strings.LastIndex(exampleVersion, ":"); idx != -1 {
								exampleVersion = exampleVersion[idx+1:]
							}
							fmt.Printf("Example: %s\n", style.Gray.Render("roverctl --force "+exampleVersion))
						}
					} else {
						fmt.Printf("No cached roverctl-web images available.\n")
						fmt.Printf("Connect to the internet to download images.\n")
					}
					return nil
				}

				// Has internet, try to pull the image
				fmt.Printf("Image not found locally, attempting to pull from registry...\n")
				ghcr := registry.AuthConfig{
					ServerAddress: "ghcr.io",
				}
				encodedAuth, err := registry.EncodeAuthConfig(ghcr)
				if err != nil {
					return fmt.Errorf("failed to encode auth: %v", err)
				}

				// Check if image exists remotely
				_, err = dc.DistributionInspect(ctx, imageRef, encodedAuth)
				if err != nil {
					fmt.Printf("No matching roverctl-web image found for roverd version %s.\n%s %s\n", version, style.Gray.Render("You can find available releases at"), style.Primary.Render("https://github.com/VU-ASE/rover/releases"))
					return nil
				}

				// Pull roverctl-web
				fmt.Printf("Pulling roverctl-web %s image...\n", style.Success.Render(version))
				out, err := dc.ImagePull(ctx, imageRef, image.PullOptions{})
				if err != nil {
					fmt.Println("Error pulling image:", err)
					return nil
				}
				defer out.Close()
				if verbose {
					_, _ = io.Copy(os.Stdout, out) // Stream pull output to console
				} else {
					_, _ = io.Copy(io.Discard, out) // Discard pull output, but still wait for it to finish
				}
			}

			//
			// Start the container(s)
			//
			proxyHost := ""
			if debugMode {
				if proxyIp != "" {
					proxyHost = proxyIp
					fmt.Printf("Binding proxy server to custom IP %s\n", proxyHost)
				} else {
					fmt.Printf("Detecting local IP address...\n")
					proxyHost, err = utils.GetLocalIP()
					if err != nil {
						fmt.Printf("Failed to detect local IP. Specify one manually with the --proxy-ip flag (%s)", err.Error())
						return nil
					}
				}
				fmt.Printf("Using local IP %s for debugging\n", proxyHost)
			}

			proxyHttpPort := 7500
			proxyUdpPort := 40000

			// conn should already be set from version determination above
			// Environment variables roverctl-web needs
			envVars := []string{
				"PUBLIC_ROVERD_HOST=" + conn.Host,
				"PUBLIC_ROVERD_PORT=80",
				"PUBLIC_ROVERD_USERNAME=" + conn.Username,
				"PUBLIC_ROVERD_PASSWORD=" + conn.Password,
				"PUBLIC_PASSTHROUGH_HOST=" + proxyHost,
				"PUBLIC_PASSTHROUGH_PORT=" + fmt.Sprintf("%d", proxyHttpPort),
			}

			webPort := 3000
			webPortStr := fmt.Sprintf("%d", webPort)
			// Port forwarding (host:container)
			portBindings := nat.PortMap{
				nat.Port(webPortStr + "/tcp"): []nat.PortBinding{
					{HostIP: "0.0.0.0", HostPort: webPortStr},
				},
			}

			// Check if all ports are available
			ports := []int{webPort}
			if debugMode {
				ports = []int{webPort, proxyHttpPort}
			}

			for _, port := range ports {
				if !utils.IsPortAvailable(port) {
					fmt.Printf("Roverctl cannot be started because port %s is in use. \nClose any applications using this port and try again.\n", style.Primary.Render(fmt.Sprintf("%d", port)))

					process, err := utils.GetProcessUsingPort(port)
					if err == nil {
						fmt.Printf("%s", style.Error.Render(process))
					}
					return nil
				}
			}

			// Create a container with the specified image and environment variables
			resp, err := dc.ContainerCreate(ctx, &container.Config{
				Image: imageRef,
				Env:   envVars,
				Tty:   true, // Keep terminal session interactive
				ExposedPorts: nat.PortSet{
					"3000/tcp": struct{}{}, // Expose container's port 3000
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
			// Check if we're in offline mode for final message
			hasInternet := checkInternetConnectivity()
			if !hasInternet {
				fmt.Printf("Visit %s to control this Rover! %s\n%s\n",
					style.Primary.Render(url),
					style.Warning.Render("(offline mode)"),
					style.Gray.Render("Press Ctrl+C to stop roverctl-web gracefully"))
			} else {
				fmt.Printf("Visit %s to control this Rover!\n%s\n",
					style.Primary.Render(url),
					style.Gray.Render("Press Ctrl+C to stop roverctl-web gracefully"))
			}
			_ = utils.OpenBrowser(url)

			// Set up signal handling (to stop container on Ctrl+C)
			sigChan := make(chan os.Signal, 1)
			signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

			go func() {
				<-sigChan // Wait for Ctrl+C (SIGINT)
				fmt.Println("\nQuitting roverctl-web" + style.Gray.Render(" (this may take a few seconds)") + "...")

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
	rootCmd.Flags().IntVarP(&roverIndex, "rover", "r", -1, "The index of the rover to upload to, between 1 and 20")
	rootCmd.Flags().StringVarP(&roverdHost, "host", "", "", "The roverd endpoint to connect to (if not using --rover)")
	rootCmd.Flags().StringVarP(&roverdUsername, "username", "u", "debix", "The username to use to connect to the roverd endpoint")
	rootCmd.Flags().StringVarP(&roverdPassword, "password", "p", "debix", "The password to use to connect to the roverd endpoint")
	rootCmd.Flags().BoolVarP(&debugMode, "debug", "d", false, "Enable debug/tuning mode")
	rootCmd.Flags().StringVarP(&proxyIp, "proxy-ip", "", "", "Override the locally detected IP address to bind the proxy server to")
	rootCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Show verbose output")
	rootCmd.Flags().StringVarP(&roverctlVersion, "force", "f", "", "Force roverctl-web to run at a specific version")

	return rootCmd
}
