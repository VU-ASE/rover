package command_pipeline

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/cobra"

	command_prechecks "github.com/VU-ASE/rover/roverctl/src/commands/prechecks"
	"github.com/VU-ASE/rover/roverctl/src/openapi"
	"github.com/VU-ASE/rover/roverctl/src/style"
	utils "github.com/VU-ASE/rover/roverctl/src/utils"
)

func Add(rootCmd *cobra.Command) {
	// General flags
	var roverIndex int
	var roverdHost string
	var roverdUsername string
	var roverdPassword string

	// pipeline command
	var infoCmd = &cobra.Command{
		Use:     "pipeline",
		Aliases: []string{"pipe", "p", "pi", "pl"},
		Short:   "Get the currently active pipeline",
		RunE: func(cmd *cobra.Command, args []string) error {
			conn, er := command_prechecks.Perform(cmd, args, roverIndex, roverdHost, roverdUsername, roverdPassword)
			if er != nil {
				return er
			}
			api := conn.ToApiClient()

			pipeline := api.PipelineAPI.PipelineGet(
				context.Background(),
			)
			res, http, err := pipeline.Execute()
			if err != nil {
				fmt.Printf("Could not fetch pipeline: %s\n", utils.ParseHTTPError(err, http))
				return nil
			}

			statusStr := style.Gray.Render("unknown")
			if res.Status == openapi.EMPTY {
				statusStr = style.Warning.Render("empty")
			} else if res.Status == openapi.STARTED {
				statusStr = style.Success.Render("started")
			} else if res.Status == openapi.STARTABLE {
				statusStr = style.Primary.Render("startable")
			}

			fmt.Printf("Pipeline status: %s\n", statusStr)
			for _, enabled := range res.Enabled {
				fmt.Println("- " + enabled.Service.Fq.Author + "/" + enabled.Service.Fq.Name + " (" + enabled.Service.Fq.Version + ")")
			}

			if res.Status == openapi.STARTED {
				fmt.Printf("Pipeline has been running since %s\n", style.Primary.Render(time.Unix(*res.LastStart/1000, 0).String()))
			}

			return nil
		},
	}
	infoCmd.Flags().IntVarP(&roverIndex, "rover", "r", 0, "The index of the rover to upload to, between 1 and 20")
	infoCmd.Flags().StringVarP(&roverdHost, "host", "", "", "The roverd endpoint to connect to (if not using --rover)")
	infoCmd.Flags().StringVarP(&roverdUsername, "username", "u", "debix", "The username to use to connect to the roverd endpoint")
	infoCmd.Flags().StringVarP(&roverdPassword, "password", "p", "debix", "The password to use to connect to the roverd endpoint")

	addStart(infoCmd)
	addStop(infoCmd)
	addReset(infoCmd)
	addEnable(infoCmd)
	addDisable(infoCmd)
	rootCmd.AddCommand(infoCmd)
}
