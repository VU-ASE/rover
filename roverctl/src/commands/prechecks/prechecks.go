package command_prechecks

import (
	"fmt"
	"strings"

	"github.com/VU-ASE/rover/roverctl/src/configuration"
	"github.com/VU-ASE/rover/roverctl/src/utils"
	view_incompatible "github.com/VU-ASE/rover/roverctl/src/views/incompatible"

	"github.com/spf13/cobra"
)

func Perform(cmd *cobra.Command, args []string, roverIndex int, roverdHost string, roverdUsername, roverdPassword string) (*configuration.RoverConnection, error) {
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

		// Pad number to two digits and use mDNS to resolve the rover's hostname
		host = fmt.Sprintf("rover%02d.local", roverIndex)
		// host = fmt.Sprintf("192.168.0.%d", roverIndex+100)

		if strings.HasSuffix(host, ".local") {
			if ip, err := utils.ResolveHostWithPing(host); err == nil {
				host = ip
			}
		}
	}

	// Create connection
	conn := configuration.RoverConnection{
		Identifier: identifier,
		Host:       host,
		Username:   roverdUsername,
		Password:   roverdPassword,
	}
	view_incompatible.WarnOnIncompatible(conn)

	return &conn, nil
}
