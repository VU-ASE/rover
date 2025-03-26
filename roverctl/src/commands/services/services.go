package command_services

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
	var lines int

	// services command
	var servicesCmd = &cobra.Command{
		Use:   "services",
		Short: "Fetch the currently installed services",
		RunE: func(cmd *cobra.Command, args []string) error {
			conn, er := command_prechecks.Perform(cmd, args, roverIndex, roverdHost, roverdUsername, roverdPassword)
			if er != nil {
				return er
			}
			api := conn.ToApiClient()

			pipeline := api.ServicesAPI.FqnsGet(
				context.Background(),
			)
			res, http, err := pipeline.Execute()
			if err != nil {
				fmt.Printf("Could not fetch services: %s\n", utils.ParseHTTPError(err, http))
				return nil
			}

			if len(res) <= 0 {
				fmt.Println("No services are currently installed")
			} else {
				for _, installed := range res {
					fmt.Println("- " + installed.Author + "/" + installed.Name + " (" + installed.Version + ")")
				}
			}

			return nil
		},
	}
	servicesCmd.Flags().IntVarP(&roverIndex, "rover", "r", 0, "The index of the rover to upload to, between 1 and 20")
	servicesCmd.Flags().StringVarP(&roverdHost, "host", "", "", "The roverd endpoint to connect to (if not using --rover)")
	servicesCmd.Flags().StringVarP(&roverdUsername, "username", "u", "debix", "The username to use to connect to the roverd endpoint")
	servicesCmd.Flags().StringVarP(&roverdPassword, "password", "p", "debix", "The password to use to connect to the roverd endpoint")
	servicesCmd.Flags().IntVarP(&lines, "lines", "l", 50, "The number of log lines to display")

	addInstall(servicesCmd)
	addDelete(servicesCmd)

	rootCmd.AddCommand(servicesCmd)
}
