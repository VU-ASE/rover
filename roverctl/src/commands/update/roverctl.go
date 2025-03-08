package command_update

import (
	"strings"

	"github.com/spf13/cobra"

	utils "github.com/VU-ASE/rover/roverctl/src/utils"
)

func addUpdateRoverctl(rootCmd *cobra.Command) {
	// Self-update command
	var version string
	var selfUpdateCmd = &cobra.Command{
		Use:   "roverctl",
		Short: "Self-update roverctl",
		Long:  `Update roverctl to the latest version, or the version specified.`,
		Run: func(cmd *cobra.Command, args []string) {
			if version != "" {
				version = strings.TrimPrefix(version, "v")
				utils.ExecuteShellCommand("curl -fsSL https://raw.githubusercontent.com/VU-ASE/rover/refs/heads/main/roverctl/install.sh | bash -s v" + version)
			} else {
				utils.ExecuteShellCommand("curl -fsSL https://raw.githubusercontent.com/VU-ASE/rover/refs/heads/main/roverctl/install.sh | bash")
			}
		},
	}
	selfUpdateCmd.Flags().StringVarP(&version, "version", "v", "", "The version tag to update/downgrade to (e.g. v0.1.0)")
	rootCmd.AddCommand(selfUpdateCmd)
}
