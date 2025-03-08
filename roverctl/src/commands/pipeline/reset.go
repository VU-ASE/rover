package command_pipeline

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	command_prechecks "github.com/VU-ASE/rover/roverctl/src/commands/prechecks"
	openapi "github.com/VU-ASE/rover/roverctl/src/openapi"
	utils "github.com/VU-ASE/rover/roverctl/src/utils"
)

func addReset(rootCmd *cobra.Command) {
	// General flags
	var roverIndex int
	var roverdHost string
	var roverdUsername string
	var roverdPassword string

	// pipeline command
	var infoCmd = &cobra.Command{
		Use:   "reset",
		Short: "Reset the currently active pipeline",
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
		},
	}
	infoCmd.Flags().IntVarP(&roverIndex, "rover", "r", 0, "The index of the rover to upload to, between 1 and 20")
	infoCmd.Flags().StringVarP(&roverdHost, "host", "", "", "The roverd endpoint to connect to (if not using --rover)")
	infoCmd.Flags().StringVarP(&roverdUsername, "username", "u", "debix", "The username to use to connect to the roverd endpoint")
	infoCmd.Flags().StringVarP(&roverdPassword, "password", "p", "debix", "The password to use to connect to the roverd endpoint")
	rootCmd.AddCommand(infoCmd)
}
