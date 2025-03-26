package command_author

import (
	"fmt"
	"regexp"

	"github.com/spf13/cobra"

	strings "strings"

	state "github.com/VU-ASE/rover/roverctl/src/state"
	style "github.com/VU-ASE/rover/roverctl/src/style"
)

func Add(rootCmd *cobra.Command) {
	// General flags
	var authorName string

	// info command
	var infoCmd = &cobra.Command{
		Use:     "author",
		Aliases: []string{"a"},
		Short:   "Set the author name used for service uploads",
		RunE: func(cmd *cobra.Command, args []string) error {
			s := state.Get()
			if s == nil {
				fmt.Printf("Roverctl state could not be recovered, please remove your configuration directory and try again\n")
				return nil
			}

			if authorName == "" {
				if s.Config.Author == "" {
					fmt.Printf("No author name set\n")
				} else {
					fmt.Printf("Current author name: %s\n", style.Primary.Render(s.Config.Author))
				}
			} else {
				// Validate
				// Can only contain lowercase letters and hyphens
				valid := regexp.MustCompile(`^[a-z0-9-]*$`).MatchString(authorName)
				if !valid {
					fmt.Printf("Cannot set author name: names can only contain lowercase letters, numbers, and hyphens\n")
				} else if len(authorName) < 3 {
					fmt.Printf("Cannot set author name: names must be at least 3 characters long\n")
				} else if strings.EqualFold(authorName, "vu-ase") {
					fmt.Printf("You cannot set your author name to 'vu-ase' (nice try buddy)\n")
				} else {
					s.Config.Author = authorName
				}
			}
			return nil
		},
	}
	infoCmd.Flags().StringVarP(&authorName, "set", "s", "", "The author name to set")
	rootCmd.AddCommand(infoCmd)
}
