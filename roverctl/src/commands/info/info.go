package command_ssh

import (
	"github.com/spf13/cobra"

	command_prechecks "github.com/VU-ASE/rover/roverctl/src/commands/prechecks"

	view_info "github.com/VU-ASE/rover/roverctl/src/views/info"
	tea "github.com/charmbracelet/bubbletea"
)

func Add(rootCmd *cobra.Command) {
	// General flags
	var roverIndex int
	var roverdHost string
	var roverdUsername string
	var roverdPassword string

	// info command
	var infoCmd = &cobra.Command{
		Use:   "info",
		Short: "View roverctl and roverd information",
		Long:  `Display build and connection information for roverctl, and roverd if a rover is specified.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			// Ignore errors
			conn, _ := command_prechecks.Perform(cmd, args, roverIndex, roverdHost, roverdUsername, roverdPassword)

			p := tea.NewProgram(view_info.New(conn))
			_, err := p.Run()
			return err
		},
	}
	infoCmd.Flags().IntVarP(&roverIndex, "rover", "r", 0, "The index of the rover to upload to, between 1 and 20")
	infoCmd.Flags().StringVarP(&roverdHost, "host", "", "", "The roverd endpoint to connect to (if not using --rover)")
	infoCmd.Flags().StringVarP(&roverdUsername, "username", "u", "debix", "The username to use to connect to the roverd endpoint")
	infoCmd.Flags().StringVarP(&roverdPassword, "password", "p", "debix", "The password to use to connect to the roverd endpoint")
	rootCmd.AddCommand(infoCmd)
}
