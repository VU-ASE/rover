package command_ssh

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"

	"github.com/spf13/cobra"

	command_prechecks "github.com/VU-ASE/rover/roverctl/src/commands/prechecks"
)

func Add(rootCmd *cobra.Command) {
	// General flags
	var roverIndex int
	var roverdHost string
	var roverdUsername string
	var roverdPassword string

	// SSH command
	var sshCmd = &cobra.Command{
		Use:   "ssh",
		Short: "Open an SSH terminal to a Rover",
		Long:  `Will use native SSH to open a terminal to the Rover.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			conn, err := command_prechecks.Perform(cmd, args, roverIndex, roverdHost, roverdUsername, roverdPassword)
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
}
