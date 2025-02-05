package main

import (
	"os"

	"github.com/VU-ASE/roverctl/src/configuration"
	"github.com/VU-ASE/roverctl/src/state"
	"github.com/VU-ASE/roverctl/src/utils"
	"github.com/VU-ASE/roverctl/src/views"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func run() error {
	// Initialize the app and create app state
	err := configuration.Initialize()
	if err != nil {
		return err
	}
	appState := state.Get()

	//
	// CLI options
	//

	// The default tui
	var rootCmd = &cobra.Command{
		Use:   "roverctl",
		Short: "Roverctl terminal user interface",
		Long:  "A terminal user interface (TUI) to manage your Rover.",
		Run: func(cmd *cobra.Command, args []string) {
			// We start the app in a separate (full) screen
			p := tea.NewProgram(views.RootScreen(appState), tea.WithAltScreen())
			_, _ = p.Run()
		},
	}

	// Entering upload mode
	var watch bool
	var uploadCmd = &cobra.Command{
		Use:   "upload <PATHS>",
		Short: "Upload specified service folders to your Rover",
		Long: `The upload command allows you to upload one or more service folders to the Rover. 
You can optionally specify the --watch flag to enable file watch and upload.`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			// Check if watch flag is enabled
			if watch {
				// We start the app in a separate (full) screen
				p := tea.NewProgram(views.CliRootScreen(appState, views.NewServicesSyncPage(args)), tea.WithAltScreen())
				_, _ = p.Run()
			} else {
				// We start the app a normal screen (output preserved)
				appState.Interactive = false
				p := tea.NewProgram(views.CliRootScreen(appState, views.NewServicesSyncPage(args)))
				_, _ = p.Run()
			}
		},
	}
	uploadCmd.Flags().BoolVarP(&watch, "watch", "w", false, "Enable file watching")
	rootCmd.AddCommand(uploadCmd)

	err = rootCmd.Execute()
	if err != nil {
		log.Err(err)
		// don't return, clean up
	}

	// Save the connections to disk
	err = state.Get().RoverConnections.Save()
	if err != nil {
		return err
	}
	err = state.Get().Config.Save()
	if err != nil {
		return err
	}

	quitCmd := state.Get().QuitCommand
	if quitCmd != "" {
		return utils.ExecuteShellCommand(quitCmd)
	}

	return nil
}

func main() {
	// Configure zerolog to output to stdout beautifully
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}) //

	// Run the program
	if err := run(); err != nil {
		log.Error().Err(err).Msg("An error occurred while running the program.")
		os.Exit(1)
	}
}
