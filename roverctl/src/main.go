package main

import (
	"os"

	"github.com/VU-ASE/rover/roverctl/src/configuration"
	"github.com/VU-ASE/rover/roverctl/src/state"
	"github.com/VU-ASE/rover/roverctl/src/utils"

	commands "github.com/VU-ASE/rover/roverctl/src/commands"
	command_author "github.com/VU-ASE/rover/roverctl/src/commands/author"
	command_calibrate "github.com/VU-ASE/rover/roverctl/src/commands/calibrate"
	command_emergency "github.com/VU-ASE/rover/roverctl/src/commands/emergency"
	command_info "github.com/VU-ASE/rover/roverctl/src/commands/info"
	command_logs "github.com/VU-ASE/rover/roverctl/src/commands/logs"
	command_pipeline "github.com/VU-ASE/rover/roverctl/src/commands/pipeline"
	command_services "github.com/VU-ASE/rover/roverctl/src/commands/services"
	command_shutdown "github.com/VU-ASE/rover/roverctl/src/commands/shutdown"

	command_ssh "github.com/VU-ASE/rover/roverctl/src/commands/ssh"
	command_update "github.com/VU-ASE/rover/roverctl/src/commands/update"
	command_upload "github.com/VU-ASE/rover/roverctl/src/commands/upload"

	//

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func run() error {
	// Initialize the app and create app state
	err := configuration.Initialize()
	if err != nil {
		return err
	}

	//
	// Build the CLI and commands
	//
	rootCmd := commands.NewRoot()
	command_pipeline.Add(rootCmd)
	command_services.Add(rootCmd)
	command_upload.Add(rootCmd)
	command_logs.Add(rootCmd)
	command_ssh.Add(rootCmd)
	command_info.Add(rootCmd)
	command_author.Add(rootCmd)
	command_update.Add(rootCmd)
	command_calibrate.Add(rootCmd)
	command_shutdown.Add(rootCmd)
	command_emergency.Add(rootCmd)

	err = rootCmd.Execute()
	if err != nil {
		log.Err(err)
		// don't return, clean up
	}

	// Save configs to disk
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
