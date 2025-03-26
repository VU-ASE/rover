package command_shutdown

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	command_prechecks "github.com/VU-ASE/rover/roverctl/src/commands/prechecks"
	"github.com/VU-ASE/rover/roverctl/src/style"
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
		Use:     "shutdown",
		Aliases: []string{"sd"},
		Short:   "Shut down the Debix immediately",
		RunE: func(cmd *cobra.Command, args []string) error {
			conn, er := command_prechecks.Perform(cmd, args, roverIndex, roverdHost, roverdUsername, roverdPassword)
			if er != nil {
				return er
			}
			api := conn.ToApiClient()

			shutdown := api.HealthAPI.ShutdownPost(
				context.Background(),
			)
			http, err := shutdown.Execute()
			if err != nil {
				fmt.Printf("Could not shut down: %s\n", utils.ParseHTTPError(err, http))
				return nil
			}

			fmt.Printf("Rover was successfully shut down\n")
			fmt.Printf("%s.\n", style.Warning.Render("When you see the unplug symbol, unplug power immediately"))
			return nil
		},
	}
	infoCmd.Flags().IntVarP(&roverIndex, "rover", "r", 0, "The index of the rover to upload to, between 1 and 20")
	infoCmd.Flags().StringVarP(&roverdHost, "host", "", "", "The roverd endpoint to connect to (if not using --rover)")
	infoCmd.Flags().StringVarP(&roverdUsername, "username", "u", "debix", "The username to use to connect to the roverd endpoint")
	infoCmd.Flags().StringVarP(&roverdPassword, "password", "p", "debix", "The password to use to connect to the roverd endpoint")

	rootCmd.AddCommand(infoCmd)
}
