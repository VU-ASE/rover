package command_upload

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	command_prechecks "github.com/VU-ASE/rover/roverctl/src/commands/prechecks"

	view_upload "github.com/VU-ASE/rover/roverctl/src/views/upload"
	tea "github.com/charmbracelet/bubbletea"
)

func Add(rootCmd *cobra.Command) {
	// General flags
	var roverIndex int
	var roverdHost string
	var roverdUsername string
	var roverdPassword string

	// Upload command
	var watch bool
	var uploadCmd = &cobra.Command{
		Use:     "upload <PATHS>",
		Aliases: []string{"u", "sync"},
		Short:   "Upload specified service folders to a Rover",
		Long: `The upload command allows you to upload one or more service folders to the Rover. 
You can optionally specify the --watch flag to enable file watch and upload.`,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return fmt.Errorf("at least one directory must be provided")
			}
			for _, dir := range args {
				info, err := os.Stat(dir)
				if err != nil {
					return fmt.Errorf("invalid directory '%s': %v", dir, err)
				}
				if !info.IsDir() {
					return fmt.Errorf("'%s' is not a directory", dir)
				}
			}
			return nil
		},

		RunE: func(cmd *cobra.Command, args []string) error {
			conn, err := command_prechecks.Perform(cmd, args, roverIndex, roverdHost, roverdUsername, roverdPassword)
			if err != nil {
				return err
			}

			p := tea.NewProgram(view_upload.New(
				*conn, args, watch,
			))
			_, err = p.Run()
			return err
		},
	}
	uploadCmd.Flags().BoolVarP(&watch, "watch", "w", false, "Enable file watching")
	uploadCmd.Flags().IntVarP(&roverIndex, "rover", "r", 0, "The index of the rover to upload to, between 1 and 20")
	uploadCmd.Flags().StringVarP(&roverdHost, "host", "", "", "The roverd endpoint to connect to (if not using --rover)")
	uploadCmd.Flags().StringVarP(&roverdUsername, "username", "u", "debix", "The username to use to connect to the roverd endpoint")
	uploadCmd.Flags().StringVarP(&roverdPassword, "password", "p", "debix", "The password to use to connect to the roverd endpoint")
	rootCmd.AddCommand(uploadCmd)
}
