package configuration

import (
	"fmt"
	"os"
)

// This is where the Rover configuration files are saved both remotely
// and on the local machine of the user.
// Directories will be created if they do not exist

func LocalConfigDir() string {
	loc := "/.config/rover"

	home, err := os.UserHomeDir()
	if err != nil {
		return "/home" + loc
	}

	return home + loc
}

// Initialize the configuration directory
func Initialize() error {
	// Check if we are root
	uid := os.Geteuid()
	if uid == 0 {
		return fmt.Errorf("You should not run this utility as root.")
	}

	// Create the configuration directory if it does not exist
	if err := os.MkdirAll(LocalConfigDir(), 0755); err != nil {
		return err
	}

	return nil
}
