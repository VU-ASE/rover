package command_pipeline

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	command_prechecks "github.com/VU-ASE/rover/roverctl/src/commands/prechecks"
	openapi "github.com/VU-ASE/rover/roverctl/src/openapi"
	utils "github.com/VU-ASE/rover/roverctl/src/utils"
)

func Add(rootCmd *cobra.Command) {
	// General flags
	var roverIndex int
	var roverdHost string
	var roverdUsername string
	var roverdPassword string
	var lines int

	// info command
	var infoCmd = &cobra.Command{
		Use:   "pipeline [start/stop/reset]",
		Short: "Get the currently active pipeline",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) == 1 {
				if args[0] != "start" && args[0] != "stop" && args[0] != "reset" {
					return fmt.Errorf("Invalid argument, specify either start, stop, reset or leave empty")
				}
			} else if len(args) > 1 {
				return fmt.Errorf("Too many arguments")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			conn, er := command_prechecks.Perform(cmd, args, roverIndex, roverdHost, roverdUsername, roverdPassword)
			if er != nil {
				return er
			}
			api := conn.ToApiClient()

			if len(args) == 0 {
				pipeline := api.PipelineAPI.PipelineGet(
					context.Background(),
				)
				res, http, err := pipeline.Execute()
				if err != nil {
					fmt.Printf("Could not fetch pipeline: %s\n", utils.ParseHTTPError(err, http))
					return nil
				}

				fmt.Printf("Pipeline status: %s\n", res.Status)
				for _, enabled := range res.Enabled {
					fmt.Println("- " + enabled.Service.Fq.Author + "/" + enabled.Service.Fq.Name + " (" + enabled.Service.Fq.Version + ")")
				}
			} else if len(args) == 1 {
				switch args[0] {
				case "start":
					{
						pipeline := api.PipelineAPI.PipelineStartPost(
							context.Background(),
						)
						http, err := pipeline.Execute()
						if err != nil {
							fmt.Printf("Could not start pipeline: %s\n", utils.ParseHTTPError(err, http))
							return nil
						}

						fmt.Printf("Pipeline started\n")
						return nil
					}
				case "stop":
					{
						pipeline := api.PipelineAPI.PipelineStopPost(
							context.Background(),
						)
						http, err := pipeline.Execute()
						if err != nil {
							fmt.Printf("Could not stop pipeline: %s\n", utils.ParseHTTPError(err, http))
							return nil
						}

						fmt.Printf("Pipeline stopped\n")
						return nil
					}
				case "reset":
					{
						// Try to stop first (best effort)
						stop := api.PipelineAPI.PipelineStopPost(
							context.Background(),
						)
						_, _ = stop.Execute()
						// Save an empty pipeline
						pipeline := api.PipelineAPI.PipelinePost(
							context.Background(),
						)
						pipeline = pipeline.PipelinePostRequestInner(
							[]openapi.PipelinePostRequestInner{},
						)

						http, err := pipeline.Execute()
						if err != nil {
							fmt.Printf("Could not reset pipeline: %s\n", utils.ParseHTTPError(err, http))
							return nil
						}

						fmt.Printf("Pipeline reset\n")
						return nil
					}
				default:
					{
						fmt.Println("Invalid argument, specify either start, stop, reset or leave empty")
					}
				}
			}

			return nil
		},
	}
	infoCmd.Flags().IntVarP(&roverIndex, "rover", "r", 0, "The index of the rover to upload to, between 1 and 20")
	infoCmd.Flags().StringVarP(&roverdHost, "host", "", "", "The roverd endpoint to connect to (if not using --rover)")
	infoCmd.Flags().StringVarP(&roverdUsername, "username", "u", "debix", "The username to use to connect to the roverd endpoint")
	infoCmd.Flags().StringVarP(&roverdPassword, "password", "p", "debix", "The password to use to connect to the roverd endpoint")
	infoCmd.Flags().IntVarP(&lines, "lines", "l", 50, "The number of log lines to display")
	rootCmd.AddCommand(infoCmd)
}
