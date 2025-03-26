package command_services

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	command_prechecks "github.com/VU-ASE/rover/roverctl/src/commands/prechecks"
	"github.com/VU-ASE/rover/roverctl/src/openapi"
	"github.com/VU-ASE/rover/roverctl/src/style"
	utils "github.com/VU-ASE/rover/roverctl/src/utils"
)

func addInstall(rootCmd *cobra.Command) {
	// General flags
	var roverIndex int
	var roverdHost string
	var roverdUsername string
	var roverdPassword string

	// services command
	var infoCmd = &cobra.Command{
		Use:   "install",
		Short: "Install a service from a given URL onto the Rover",
		RunE: func(cmd *cobra.Command, args []string) error {
			conn, er := command_prechecks.Perform(cmd, args, roverIndex, roverdHost, roverdUsername, roverdPassword)
			if er != nil {
				return er
			}
			api := conn.ToApiClient()
			if len(args) != 1 {
				return fmt.Errorf("Specify the ZIP URL of the service to install")
			}

			fmt.Printf("Attempting to install service on Rover...\n")
			fetch := api.ServicesAPI.FetchPost(
				context.Background(),
			)
			fetch = fetch.FetchPostRequest(openapi.FetchPostRequest{
				Url: args[0],
			})
			res, http, err := fetch.Execute()
			if err != nil {
				fmt.Printf("Could not install service: %s\n", utils.ParseHTTPError(err, http))
				return nil
			}

			fmt.Printf("Installed service %s by %s (%s)\n", style.Primary.Render(res.Fq.Name), style.Primary.Render(res.Fq.Author), (res.Fq.Version))
			return nil
		},
	}
	infoCmd.Flags().IntVarP(&roverIndex, "rover", "r", 0, "The index of the rover to upload to, between 1 and 20")
	infoCmd.Flags().StringVarP(&roverdHost, "host", "", "", "The roverd endpoint to connect to (if not using --rover)")
	infoCmd.Flags().StringVarP(&roverdUsername, "username", "u", "debix", "The username to use to connect to the roverd endpoint")
	infoCmd.Flags().StringVarP(&roverdPassword, "password", "p", "debix", "The password to use to connect to the roverd endpoint")

	rootCmd.AddCommand(infoCmd)
}
