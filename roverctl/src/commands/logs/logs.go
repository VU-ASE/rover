package command_logs

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	strings "strings"

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

	// logs command
	var infoCmd = &cobra.Command{
		Use:   "logs <author> <name> <version>",
		Short: "View logs for a fully qualified service",
		Long:  `View a specified number of log lines from a service fully qualified by its author, name and version`,
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

			if len(args) != 3 {
				return nil // already errored
			}

			api := conn.ToApiClient()

			author := args[0]
			name := args[1]
			version := args[2]
			version = strings.TrimPrefix(version, "v")

			logs := api.PipelineAPI.LogsAuthorNameVersionGet(
				context.Background(),
				author,
				name,
				version,
			)
			logs = logs.Lines(int32(lines))

			res, http, err := logs.Execute()
			if err != nil {
				fmt.Printf("Could not fetch logs: %s\n", utils.ParseHTTPError(err, http))
				return nil
			}

			for _, line := range res {
				fmt.Println(line)
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
