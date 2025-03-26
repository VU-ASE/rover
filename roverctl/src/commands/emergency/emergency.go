package command_emergency

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	command_prechecks "github.com/VU-ASE/rover/roverctl/src/commands/prechecks"
	utils "github.com/VU-ASE/rover/roverctl/src/utils"
)

func Add(rootCmd *cobra.Command) {
	// General flags
	var roverIndex int
	var roverdHost string
	var roverdUsername string
	var roverdPassword string

	// services command
	var infoCmd = &cobra.Command{
		Use:     "emergency",
		Aliases: []string{"e", "em", "panic", "reset"},
		Short:   "Stop any running pipeline and reset the Rover",
		RunE: func(cmd *cobra.Command, args []string) error {
			conn, er := command_prechecks.Perform(cmd, args, roverIndex, roverdHost, roverdUsername, roverdPassword)
			if er != nil {
				return er
			}
			api := conn.ToApiClient()

			emergency := api.HealthAPI.EmergencyPost(
				context.Background(),
			)
			http, err := emergency.Execute()
			if err != nil {
				fmt.Printf("Could not run emergency command: %s\n", utils.ParseHTTPError(err, http))
				return nil
			}

			fmt.Printf("Emergency reset successful\n")
			return nil
		},
	}
	infoCmd.Flags().IntVarP(&roverIndex, "rover", "r", 0, "The index of the rover to upload to, between 1 and 20")
	infoCmd.Flags().StringVarP(&roverdHost, "host", "", "", "The roverd endpoint to connect to (if not using --rover)")
	infoCmd.Flags().StringVarP(&roverdUsername, "username", "u", "debix", "The username to use to connect to the roverd endpoint")
	infoCmd.Flags().StringVarP(&roverdPassword, "password", "p", "debix", "The password to use to connect to the roverd endpoint")

	rootCmd.AddCommand(infoCmd)
}
