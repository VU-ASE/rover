package command_services

import (
	"context"
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	command_prechecks "github.com/VU-ASE/rover/roverctl/src/commands/prechecks"
	"github.com/VU-ASE/rover/roverctl/src/style"
	utils "github.com/VU-ASE/rover/roverctl/src/utils"
)

func addDelete(rootCmd *cobra.Command) {
	// General flags
	var roverIndex int
	var roverdHost string
	var roverdUsername string
	var roverdPassword string

	// services command
	var infoCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete a fully qualified service from the Rover",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 3 {
				return fmt.Errorf("exactly one fully qualified service must be provided in the form <author> <name> <version>")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			conn, er := command_prechecks.Perform(cmd, args, roverIndex, roverdHost, roverdUsername, roverdPassword)
			if er != nil {
				return er
			}
			api := conn.ToApiClient()

			author := args[0]
			name := args[1]
			version := strings.TrimPrefix(args[2], "v") // should pass "1.0.0" instead of "v1.0.0"

			delete := api.ServicesAPI.ServicesAuthorServiceVersionDelete(
				context.Background(),
				author,
				name,
				version,
			)
			_, http, err := delete.Execute()
			if err != nil {
				fmt.Printf("Could not delete service: %s\n", utils.ParseHTTPError(err, http))
				return nil
			}

			fmt.Printf("Deleted %s by %s (%s)\n", style.Primary.Render(name), style.Primary.Render(author), (version))
			return nil
		},
	}
	infoCmd.Flags().IntVarP(&roverIndex, "rover", "r", 0, "The index of the rover to upload to, between 1 and 20")
	infoCmd.Flags().StringVarP(&roverdHost, "host", "", "", "The roverd endpoint to connect to (if not using --rover)")
	infoCmd.Flags().StringVarP(&roverdUsername, "username", "u", "debix", "The username to use to connect to the roverd endpoint")
	infoCmd.Flags().StringVarP(&roverdPassword, "password", "p", "debix", "The password to use to connect to the roverd endpoint")

	rootCmd.AddCommand(infoCmd)
}
