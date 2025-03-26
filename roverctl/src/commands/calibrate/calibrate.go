package command_calibrate

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/spf13/cobra"
	"golang.org/x/term"

	command_prechecks "github.com/VU-ASE/rover/roverctl/src/commands/prechecks"
	"github.com/VU-ASE/rover/roverctl/src/configuration"
	"github.com/VU-ASE/rover/roverctl/src/openapi"
	"github.com/VU-ASE/rover/roverctl/src/style"
	utils "github.com/VU-ASE/rover/roverctl/src/utils"
)

// See https://github.com/VU-ASE/actuator-tester/blob/main/docs/01-overview.md and https://github.com/VU-ASE/actuator-tester
type actuatorTesterPayload struct {
	Channel int     `json:"channel"`
	Value   float32 `json:"value"`
}

func Add(rootCmd *cobra.Command) {
	// General flags
	var roverIndex int
	var roverdHost string
	var roverdUsername string
	var roverdPassword string

	// pipeline command
	var infoCmd = &cobra.Command{
		Use:     "calibrate",
		Aliases: []string{"cal", "c", "ca", "cali", "trim"},
		Short:   "Calibrate an actuator service",
		RunE: func(cmd *cobra.Command, args []string) error {
			conn, er := command_prechecks.Perform(cmd, args, roverIndex, roverdHost, roverdUsername, roverdPassword)
			if er != nil {
				return er
			}
			api := conn.ToApiClient()

			fmt.Printf("%s\n", style.Gray.Render("Preparation"))
			fmt.Println("Stopping currently active pipeline...")
			// Try to stop any running pipeline first (best effort)
			stop := api.PipelineAPI.PipelineStopPost(
				context.Background(),
			)
			_, _ = stop.Execute()

			// This pipeline should be restored later
			fmt.Println("Fetching configured pipeline...")
			pipeline := api.PipelineAPI.PipelineGet(
				context.Background(),
			)
			res, http, err := pipeline.Execute()
			if err != nil {
				fmt.Printf("Could not fetch pipeline: %s\n", utils.ParseHTTPError(err, http))
				return nil
			}

			// Get all services, see if the actuator and actuator-tester is already installed
			fqns := api.ServicesAPI.FqnsGet(
				context.Background(),
			)
			fqnsRes, http, err := fqns.Execute()
			if err != nil {
				fmt.Printf("Could not fetch services: %s\n", utils.ParseHTTPError(err, http))
				return nil
			}

			actuators := []openapi.FullyQualifiedService{}
			actuatorTester := []openapi.FullyQualifiedService{}
			for _, fqn := range fqnsRes {
				if fqn.Name == "actuator" {
					actuators = append(actuators, fqn)
				} else if fqn.Name == "actuator-tester" {
					actuatorTester = append(actuatorTester, fqn)
				}
			}

			fmt.Printf("\n%s\n", style.Gray.Render("Actuator target"))
			// Try to find the actuator service, or install if missing
			var actuator openapi.FullyQualifiedService
			if len(actuators) == 0 {
				// Install actuator
				fmt.Println("No actuator service found. Installing it...")
				// Get the latest version
				actuatorAvailable, err := utils.CheckForGithubUpdate("actuator", configuration.VU_ASE_AUTHOR, "")
				if err != nil {
					fmt.Printf("Could not find actuator to install %s\n", err)
					return nil
				} else if actuatorAvailable == nil {
					fmt.Println("No actuator available to install")
					return nil
				}
				fmt.Printf("Installing actuator %s...\n", actuatorAvailable.LatestVersion)
				var asset utils.UpdateAvailableAsset
				// Find the asset with the zip extension
				for _, a := range actuatorAvailable.Assets {
					if a.ContentType == "application/zip" {
						asset = a
						break
					}
				}
				if asset.Url == "" {
					fmt.Printf("No assets found for actuator %s. The release might still be building (check back in 10 minutes!). Otherwise, try to install the actuator service manually.\n", actuatorAvailable.LatestVersion)
					return nil
				}

				fetchActuator := api.ServicesAPI.FetchPost(
					context.Background(),
				)
				fetchActuator = fetchActuator.FetchPostRequest(
					openapi.FetchPostRequest{
						Url: asset.Url,
					},
				)
				fetchRes, http, err := fetchActuator.Execute()
				if err != nil {
					fmt.Printf("Could not install actuator: %s\n", utils.ParseHTTPError(err, http))
					return nil
				}
				fmt.Printf("Installed %s by %s (%s)\n", fetchRes.Fq.Name, fetchRes.Fq.Author, fetchRes.Fq.Version)
				actuator = fetchRes.Fq
			} else if len(actuators) == 1 {
				actuator = actuators[0]
			} else {
				fmt.Println("Multiple actuators found")
				for i, a := range actuators {
					fmt.Printf("[%d]: %s by %s (%s)\n", i, a.Name, style.Primary.Render(a.Author), a.Version)
				}
				i := -1
				var err error
				for i < 0 || i >= len(actuators) || err != nil {
					fmt.Printf("Select the actuator to use by specifying an index between 0 and %d\n >", len(actuators)-1)
					_, err = fmt.Scan(&i)
				}
				actuator = actuators[i]
			}
			fmt.Printf("Selected calibration target: %s by %s (%s)\n", style.Primary.Render(actuator.Name), style.Primary.Render(actuator.Author), actuator.Version)

			fmt.Printf("\n%s\n", style.Gray.Render("Actuator-tester tooling"))
			// Try to find the actuator-tester service, or install if missing
			var actuatorTesterService openapi.FullyQualifiedService
			if len(actuatorTester) == 0 {
				// Install actuator-tester
				fmt.Println("No actuator-tester service found. Installing it...")
				// Get the latest version
				actuatorTesterAvailable, err := utils.CheckForGithubUpdate("actuator-tester", configuration.VU_ASE_AUTHOR, "")
				if err != nil {
					fmt.Printf("Could not find actuator-tester to install %s\n", err)
					return nil
				} else if actuatorTesterAvailable == nil {
					fmt.Println("No actuator-tester available to install")
					return nil
				}
				fmt.Printf("Installing actuator-tester %s...\n", actuatorTesterAvailable.LatestVersion)
				var asset utils.UpdateAvailableAsset
				// Find the asset with the zip extension
				for _, a := range actuatorTesterAvailable.Assets {
					if a.ContentType == "application/zip" {
						asset = a
						break
					}
				}
				if asset.Url == "" {
					fmt.Printf("No assets found for actuator-tester %s. The release might still be building (check back in 10 minutes!). Otherwise, try to install the actuator-tester service manually.\n", actuatorTesterAvailable.LatestVersion)
					return nil
				}

				fetchActuatorTester := api.ServicesAPI.FetchPost(
					context.Background(),
				)
				fetchActuatorTester = fetchActuatorTester.FetchPostRequest(
					openapi.FetchPostRequest{
						Url: asset.Url,
					},
				)
				fetchRes, http, err := fetchActuatorTester.Execute()
				if err != nil {
					fmt.Printf("Could not install actuator-tester: %s\n", utils.ParseHTTPError(err, http))
					return nil
				}
				fmt.Printf("Installed %s by %s (%s)\n", fetchRes.Fq.Name, fetchRes.Fq.Author, fetchRes.Fq.Version)
				actuatorTesterService = fetchRes.Fq
			} else if len(actuatorTester) == 1 {
				actuatorTesterService = actuatorTester[0]
			} else {
				for i, a := range actuatorTester {
					fmt.Printf("[%d]: %s by %s (%s)\n", i, a.Name, style.Primary.Render(a.Author), a.Version)
				}
				i := -1
				var err error
				for i < 0 || i >= len(actuatorTester) || err != nil {
					fmt.Printf("Select the actuator-tester to use by specifying an index between 0 and %d\n> ", len(actuatorTester)-1)
					_, err = fmt.Scan(&i)
				}
				actuatorTesterService = actuatorTester[i]
			}
			fmt.Printf("Selected tooling: %s by %s (%s)\n", style.Primary.Render(actuatorTesterService.Name), style.Primary.Render(actuatorTesterService.Author), actuatorTesterService.Version)

			// Set the UDP port for the actuator tester to listen on
			fmt.Printf("\n%s\n", style.Gray.Render("Calibration configuration"))
			actuatorTesterPort := 12345
			actuatorTesterPortStr := fmt.Sprintf(":%d", actuatorTesterPort) // need this to create a pointer for the API call, not great
			actuatorTesterPortRequest := api.ServicesAPI.ServicesAuthorServiceVersionConfigurationPost(
				context.Background(),
				actuatorTesterService.Author,
				actuatorTesterService.Name,
				actuatorTesterService.Version,
			)
			actuatorTesterPortRequest = actuatorTesterPortRequest.ServicesAuthorServiceVersionConfigurationPostRequestInner(
				[]openapi.ServicesAuthorServiceVersionConfigurationPostRequestInner{
					{
						Key: "udp-port", // see: https://github.com/VU-ASE/actuator-tester/blob/f713492da360f00dfb9537240e82db5534a8f3bc/service.yaml#L18
						Value: openapi.ServicesAuthorServiceVersionConfigurationPostRequestInnerValue{
							String: &actuatorTesterPortStr,
						},
					},
				},
			)
			http, err = actuatorTesterPortRequest.Execute()
			if err != nil {
				fmt.Printf("Could not set UDP port to %d for actuator-tester: %s\n", actuatorTesterPort, utils.ParseHTTPError(err, http))
				return nil
			}
			fmt.Printf("Set UDP port to %d for actuator-tester\n", actuatorTesterPort)

			//
			// At this point, we have:
			// - an actuator service installed and selected for trimming
			// - the actuator-tester service installed and configured to listen on a UDP port in our control
			// now, set up the actuator-tester --> actuator pipeline to start calibrating the actuators
			//

			// Create new pipeline
			fmt.Println("Setting calibration pipeline...")
			calibrationPipeline := api.PipelineAPI.PipelinePost(
				context.Background(),
			)
			calibrationPipeline = calibrationPipeline.PipelinePostRequestInner([]openapi.PipelinePostRequestInner{
				{
					Fq: actuator,
				},
				{
					Fq: actuatorTesterService,
				},
			})
			http, err = calibrationPipeline.Execute()
			if err != nil {
				fmt.Printf("Could not save calibration pipeline: %s\n", utils.ParseHTTPError(err, http))
				return nil
			}
			// Start the pipeline
			fmt.Println("Starting calibration pipeline...")
			start := api.PipelineAPI.PipelineStartPost(
				context.Background(),
			)
			http, err = start.Execute()
			if err != nil {
				fmt.Printf("Could not start calibration pipeline: %s\n", utils.ParseHTTPError(err, http))
				return nil
			}

			//
			// calibration pipeline is running!
			// start sending over UDP
			//

			// Create UDP address
			addr := net.JoinHostPort(conn.Host, fmt.Sprintf("%d", actuatorTesterPort))
			udpAddr, err := net.ResolveUDPAddr("udp", addr)
			if err != nil {
				fmt.Println("Error resolving UDP address:", err)
				return nil
			}
			// Dial UDP
			testerConn, err := net.DialUDP("udp", nil, udpAddr)
			if err != nil {
				fmt.Println("Error dialing UDP:", err)
				return nil
			}
			defer testerConn.Close()

			fmt.Printf("\nNow in %s!\n", style.Primary.Bold(true).Render("calibration mode"))
			fmt.Printf("%s\n", style.Success.Render("Take a look at the Rover and adjust the trim value to make it steer straight"))
			fmt.Printf("- use %s and %s to steer left and right\n", style.Primary.Render("←"), style.Primary.Render("→"))
			fmt.Printf("- use %s and %s to adjust the trim delta precision\n", style.Primary.Render("↑"), style.Primary.Render("↓"))
			fmt.Printf("- press %s to save and quit\n\n", style.Primary.Render("q"))
			// Save the current terminal state to restore later
			oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
			if err != nil {
				panic(err)
			}

			var buf [3]byte
			trimValue := float32(0.0)  // float32 because this is how the Configuration POST API expects it
			trimDelta := float32(0.01) // by how much to change the trim value

			for {
				os.Stdin.Read(buf[:])

				switch {
				case buf[0] == 0x1b && buf[1] == 0x5b && buf[2] == 0x44:
					fmt.Print("\r← ")
					trimValue -= trimDelta
				case buf[0] == 0x1b && buf[1] == 0x5b && buf[2] == 0x43:
					fmt.Print("\r→ ")
					trimValue += trimDelta

				case buf[0] == 0x1b && buf[1] == 0x5b && buf[2] == 0x41:
					fmt.Print("\r↑ ")
					trimDelta *= 10
				case buf[0] == 0x1b && buf[1] == 0x5b && buf[2] == 0x42:
					fmt.Print("\r↓ ")
					trimDelta /= 10

				case buf[0] == 'q':
					fmt.Println("\rQuitting calibration mode")
				default:
					fmt.Printf("\rUnknown input: %v", buf)
				}
				if buf[0] == 'q' {
					break
				}
				if trimValue < -1.0 {
					trimValue = -1.0
				} else if trimValue > 1.0 {
					trimValue = 1.0
				}
				fmt.Printf("Trim: %.2f ", trimValue)
				fmt.Printf(" precision: %.2f", trimDelta)

				// Check if the calibration pipeline is still running
				pipelineStatus := api.PipelineAPI.PipelineGet(
					context.Background(),
				)
				pipeline, http, err := pipelineStatus.Execute()
				if err != nil {
					fmt.Printf("\nCould not fetch pipeline status: %s\n", utils.ParseHTTPError(err, http))
					return nil
				}
				if pipeline.Status != openapi.STARTED {
					fmt.Println("\nCalibration pipeline is not running anymore. Was it stopped by someone else?")
					break
				} else if len(pipeline.Enabled) != 2 {
					fmt.Println("\nCalibration pipeline is running, but not all services are enabled. Did someone change the pipeline?")
					break
				}
				// Try to find both services
				actuatorEnabled := false
				actuatorTesterEnabled := false
				for _, enabled := range pipeline.Enabled {
					if enabled.Service.Fq.Name == actuator.Name {
						actuatorEnabled = true
					} else if enabled.Service.Fq.Name == actuatorTesterService.Name {
						actuatorTesterEnabled = true
					}
				}
				if !actuatorEnabled {
					fmt.Println("\nCalibration pipeline is running, but the actuator service is not enabled. Was it stopped?")
					break
				} else if !actuatorTesterEnabled {
					fmt.Println("\nCalibration pipeline is running, but the actuator-tester service is not enabled. Was it stopped?")
					break
				}

				//
				// Every calibration is a two-step:
				// send a full-left message (otherwise the servo does not steer as much)
				// then send the trim value
				//

				// Create message
				msg := actuatorTesterPayload{
					Channel: 0,
					Value:   1.0,
				}
				// Marshal to JSON
				jsonData, err := json.Marshal(msg)
				if err != nil {
					fmt.Println("\nError marshaling JSON:", err)
					return nil
				}
				// Send it
				_, err = testerConn.Write(jsonData)
				if err != nil {
					fmt.Println("\nError sending UDP message:", err)
					return nil
				}
				time.Sleep(100 * time.Millisecond)
				// Create message
				msg = actuatorTesterPayload{
					Channel: 0,
					Value:   trimValue,
				}
				// Marshal to JSON
				jsonData, err = json.Marshal(msg)
				if err != nil {
					fmt.Println("\nError marshaling JSON:", err)
					return nil
				}
				// Send it
				_, err = testerConn.Write(jsonData)
				if err != nil {
					fmt.Println("\nError sending UDP message:", err)
					return nil
				}
			}
			term.Restore(int(os.Stdin.Fd()), oldState)

			// Stop calibration pipeline
			fmt.Println("\nStopping calibration pipeline...")
			// Try to stop any running pipeline first (best effort)
			stop = api.PipelineAPI.PipelineStopPost(
				context.Background(),
			)
			_, _ = stop.Execute()

			// Save the trim value to the actuator service
			fmt.Printf("Saving trim value %.2f to actuator...\n", trimValue)
			trimRequest := api.ServicesAPI.ServicesAuthorServiceVersionConfigurationPost(
				context.Background(),
				actuator.Author,
				actuator.Name,
				actuator.Version,
			)
			trimRequest = trimRequest.ServicesAuthorServiceVersionConfigurationPostRequestInner(
				[]openapi.ServicesAuthorServiceVersionConfigurationPostRequestInner{
					{
						Key: "servo-trim", // see https://github.com/VU-ASE/actuator/blob/f14947a3fdef53f196d05f863035117199cee2a1/service.yaml#L39
						Value: openapi.ServicesAuthorServiceVersionConfigurationPostRequestInnerValue{
							Float32: &trimValue,
						},
					},
				},
			)
			http, err = trimRequest.Execute()
			if err != nil {
				fmt.Printf("Could not save trim value to actuator: %s\n", utils.ParseHTTPError(err, http))
				return nil
			}

			// Restore old pipeline
			fmt.Println("Restoring configured pipeline...")
			pipelineRestore := api.PipelineAPI.PipelinePost(
				context.Background(),
			)
			restoredServices := make([]openapi.PipelinePostRequestInner, 0)
			for _, enabled := range res.Enabled {
				restoredServices = append(restoredServices, openapi.PipelinePostRequestInner{
					Fq: enabled.Service.Fq,
				})
			}
			pipelineRestore = pipelineRestore.PipelinePostRequestInner(restoredServices)
			http, err = pipelineRestore.Execute()
			if err != nil {
				fmt.Printf("Could not restore active pipeline: %s\n", utils.ParseHTTPError(err, http))
				return nil
			}

			return nil
		},
	}
	infoCmd.Flags().IntVarP(&roverIndex, "rover", "r", 0, "The index of the rover to upload to, between 1 and 20")
	infoCmd.Flags().StringVarP(&roverdHost, "host", "", "", "The roverd endpoint to connect to (if not using --rover)")
	infoCmd.Flags().StringVarP(&roverdUsername, "username", "u", "debix", "The username to use to connect to the roverd endpoint")
	infoCmd.Flags().StringVarP(&roverdPassword, "password", "p", "debix", "The password to use to connect to the roverd endpoint")
	rootCmd.AddCommand(infoCmd)
}
