package command_update

import (
	"github.com/spf13/cobra"

	"github.com/VU-ASE/rover/roverctl/src/configuration"
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
				version = utils.Version(version)
				utils.ExecuteShellCommand(configuration.ROVERCTL_UPDATE_LATEST_SCRIPT_WITH_VERSION + version)
			} else {
				utils.ExecuteShellCommand(configuration.ROVERCTL_UPDATE_LATEST_SCRIPT)
			}
		},
	}
	selfUpdateCmd.Flags().StringVarP(&version, "version", "v", "", "The version tag to update/downgrade to (e.g. v0.1.0)")
	rootCmd.AddCommand(selfUpdateCmd)
}
