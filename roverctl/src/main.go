package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"

	"github.com/VU-ASE/rover/roverctl/src/configuration"
	"github.com/VU-ASE/rover/roverctl/src/state"
	"github.com/VU-ASE/rover/roverctl/src/utils"
	view_info "github.com/VU-ASE/rover/roverctl/src/views/info"
	view_upload "github.com/VU-ASE/rover/roverctl/src/views/upload"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func prechecks(cmd *cobra.Command, args []string, roverIndex int, roverdHost string, roverdUsername, roverdPassword string) (*configuration.RoverConnection, error) {
	// XOR Logic: Exactly one must be set
	roverSet := cmd.Flags().Changed("rover")
	hostSet := cmd.Flags().Changed("host")
	if roverSet == hostSet { // both false or both true
		return nil, fmt.Errorf("you must provide either --rover or --host, but not both")
	}

	identifier := roverdHost
	host := roverdHost
	if roverSet {
		if roverIndex < 1 || roverIndex > 20 {
			return nil, fmt.Errorf("rover index must be between 1 and 20")
		}
		identifier = fmt.Sprintf("rover %d", roverIndex)
		host = fmt.Sprintf("192.168.0.%d", roverIndex+100)
	}

	// Create connection
	conn := configuration.RoverConnection{
		Identifier: identifier,
		Host:       host,
		Username:   roverdUsername,
		Password:   roverdPassword,
	}
	return &conn, nil
}

func run() error {
	// Initialize the app and create app state
	err := configuration.Initialize()
	if err != nil {
		return err
	}

	//
	// CLI options
	//

	var rootCmd = &cobra.Command{
		Use:   "roverctl",
		Short: "Roverctl terminal user interface",
		Long:  "A command line interface to manage your Rover",
		Run: func(cmd *cobra.Command, args []string) {
			// // We start the app in a separate (full) screen
			// p := tea.NewProgram(views.RootScreen(appState), tea.WithAltScreen())
			// _, _ = p.Run()

			fmt.Printf("Hello, Rover!\n")
		},
	}

	// General flags
	var watch bool
	var roverIndex int
	var roverdHost string
	var roverdUsername string
	var roverdPassword string

	// Upload command
	var uploadCmd = &cobra.Command{
		Use:   "upload <PATHS>",
		Short: "Upload specified service folders to a Rover",
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
			conn, err := prechecks(cmd, args, roverIndex, roverdHost, roverdUsername, roverdPassword)
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

	// SSH command
	var sshCmd = &cobra.Command{
		Use:   "ssh",
		Short: "Open an SSH terminal to a Rover",
		Long:  `Will use native SSH to open a terminal to the Rover.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			conn, err := prechecks(cmd, args, roverIndex, roverdHost, roverdUsername, roverdPassword)
			if err != nil {
				return err
			}

			sshPath, err := exec.LookPath("ssh") // Find the SSH binary
			if err != nil {
				fmt.Println("Could not find ssh in PATH")
			}

			host := fmt.Sprintf("%s@%s", conn.Username, conn.Host)
			err = syscall.Exec(sshPath, []string{"ssh", host}, os.Environ())
			if err != nil {
				return err
			}
			return nil
		},
	}
	sshCmd.Flags().IntVarP(&roverIndex, "rover", "r", 0, "The index of the rover to upload to, between 1 and 20")
	sshCmd.Flags().StringVarP(&roverdHost, "host", "", "", "The roverd endpoint to connect to (if not using --rover)")
	sshCmd.Flags().StringVarP(&roverdUsername, "username", "u", "debix", "The username to use to connect to the roverd endpoint")
	sshCmd.Flags().StringVarP(&roverdPassword, "password", "p", "debix", "The password to use to connect to the roverd endpoint")
	rootCmd.AddCommand(sshCmd)

	// info command
	var infoCmd = &cobra.Command{
		Use:   "info",
		Short: "View roverctl and roverd information",
		Long:  `Display build and connection information for roverctl, and roverd if a rover is specified.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			// Ignore errors
			conn, _ := prechecks(cmd, args, roverIndex, roverdHost, roverdUsername, roverdPassword)

			p := tea.NewProgram(view_info.New(conn))
			_, err = p.Run()
			return err
		},
	}
	infoCmd.Flags().IntVarP(&roverIndex, "rover", "r", 0, "The index of the rover to upload to, between 1 and 20")
	infoCmd.Flags().StringVarP(&roverdHost, "host", "", "", "The roverd endpoint to connect to (if not using --rover)")
	infoCmd.Flags().StringVarP(&roverdUsername, "username", "u", "debix", "The username to use to connect to the roverd endpoint")
	infoCmd.Flags().StringVarP(&roverdPassword, "password", "p", "debix", "The password to use to connect to the roverd endpoint")
	rootCmd.AddCommand(infoCmd)

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
