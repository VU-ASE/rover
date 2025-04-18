package command_pipeline

import (
	"context"
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	command_prechecks "github.com/VU-ASE/rover/roverctl/src/commands/prechecks"
	openapi "github.com/VU-ASE/rover/roverctl/src/openapi"
	utils "github.com/VU-ASE/rover/roverctl/src/utils"
)

func addDisable(rootCmd *cobra.Command) {
	// General flags
	var roverIndex int
	var roverdHost string
	var roverdUsername string
	var roverdPassword string

	// pipeline command
	var infoCmd = &cobra.Command{
		Use:     "disable <author> <name> [<version>]",
		Aliases: []string{"d"},
		Short:   "Disable a service in the pipeline",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 2 {
				return fmt.Errorf("exactly one fully qualified service must be provided in the form <author> <name> [<version>]")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			conn, er := command_prechecks.Perform(cmd, args, roverIndex, roverdHost, roverdUsername, roverdPassword)
			if er != nil {
				return er
			}
			api := conn.ToApiClient()

			// Try to stop first (best effort)
			stop := api.PipelineAPI.PipelineStopPost(
				context.Background(),
			)
			_, _ = stop.Execute()

			// Get the current pipeline
			pipeline := api.PipelineAPI.PipelineGet(
				context.Background(),
			)
			res, http, err := pipeline.Execute()
			if err != nil {
				fmt.Printf("Could not fetch pipeline: %s\n", utils.ParseHTTPError(err, http))
				return nil
			}

			author := args[0]
			name := args[1]

			version := ""

			if len(args) > 2 {
				version := args[2]
				version = strings.TrimPrefix(version, "v")
			}

			// Add the new service to the pipeline
			found := false
			newPipeline := []openapi.PipelinePostRequestInner{}
			for _, enabled := range res.Enabled {
				if enabled.Service.Fq.Author != author && enabled.Service.Fq.Name != name && (version == "" || enabled.Service.Fq.Version != version) {
					newPipeline = append(newPipeline,
						openapi.PipelinePostRequestInner{
							Fq: enabled.Service.Fq,
						},
					)
				} else {
					found = true
				}
			}
			// Add the new service if it wasn't already in the pipeline
			if !found {
				fmt.Printf("This service is not enabled\n")
				return nil
			}

			// Save the new pipeline
			savedPipeline := api.PipelineAPI.PipelinePost(
				context.Background(),
			)
			savedPipeline = savedPipeline.PipelinePostRequestInner(
				newPipeline,
			)

			http, err = savedPipeline.Execute()
			if err != nil {
				fmt.Printf("Could not save pipeline: %s\n", utils.ParseHTTPError(err, http))
				return nil
			}

			fmt.Printf("Service disabled\n")
			return nil
		},
	}
	infoCmd.Flags().IntVarP(&roverIndex, "rover", "r", 0, "The index of the rover to upload to, between 1 and 20")
	infoCmd.Flags().StringVarP(&roverdHost, "host", "", "", "The roverd endpoint to connect to (if not using --rover)")
	infoCmd.Flags().StringVarP(&roverdUsername, "username", "u", "debix", "The username to use to connect to the roverd endpoint")
	infoCmd.Flags().StringVarP(&roverdPassword, "password", "p", "debix", "The password to use to connect to the roverd endpoint")
	rootCmd.AddCommand(infoCmd)
}
