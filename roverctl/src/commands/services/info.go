package command_services

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/spf13/cobra"

	command_prechecks "github.com/VU-ASE/rover/roverctl/src/commands/prechecks"
	"github.com/VU-ASE/rover/roverctl/src/style"
	utils "github.com/VU-ASE/rover/roverctl/src/utils"
)

func addInfo(rootCmd *cobra.Command) {
	// General flags
	var roverIndex int
	var roverdHost string
	var roverdUsername string
	var roverdPassword string

	// services command
	var infoCmd = &cobra.Command{
		Use:     "info",
		Aliases: []string{"i", "describe", "about"},
		Short:   "View info about an installed service on the Rover",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 3 {
				return fmt.Errorf("exactly one fully qualified service must be provided in the form <author> <name> <version>")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			conn, er := command_prechecks.Perform(cmd, args, roverIndex, roverdHost, roverdUsername, roverdPassword)
			if er != nil {
				return er
			}
			api := conn.ToApiClient()

			author := args[0]
			name := args[1]
			version := strings.TrimPrefix(args[2], "v") // should pass "1.0.0" instead of "v1.0.0"

			service := api.ServicesAPI.ServicesAuthorServiceVersionGet(
				context.Background(),
				author,
				name,
				version,
			)
			res, http, err := service.Execute()
			if err != nil {
				fmt.Printf("Could not get information about service: %s\n", utils.ParseHTTPError(err, http))
				return nil
			}

			fmt.Printf("About %s by %s (%s)\n", style.Primary.Render(name), style.Primary.Render(author), (version))
			builtAt := res.BuiltAt
			if builtAt == nil {
				fmt.Printf("%s\n", style.Gray.Render("This service has not been built yet"))
			} else {
				fmt.Printf("%s\n", ("This service was last built at ")+style.Primary.Render(time.Unix(*res.BuiltAt/1000, 0).String()))
			}
			fmt.Printf("Installation directory: %s\n", style.Primary.Render("/home/debix/.rover/"+author+"/"+name+"/"+version))

			if len(res.Configuration) > 0 {
				fmt.Printf("%s\n", style.Gray.Render("\nConfigurable options:"))
				for _, c := range res.Configuration {
					floatVal := c.Value.Float32
					strVal := c.Value.String

					prefix := ""
					if c.Tunable {
						prefix = style.Success.Render(" tunable")
					}

					fmt.Printf("  -%s %s %s: ", prefix, style.Gray.Render(c.Type), (c.Name))
					if floatVal != nil {
						// Print float value with 3 decimal places
						fmt.Printf("%s", style.Primary.Render(fmt.Sprintf("%.3f", *floatVal)))
					} else if strVal != nil {
						fmt.Printf("%s", style.Primary.Render(*strVal))
					} else {
						fmt.Printf("%s", style.Primary.Render("null"))
					}
					fmt.Printf("\n")
				}
			}

			if len(res.Inputs) > 0 {
				fmt.Printf("%s\n", style.Gray.Render("\nInputs:"))
				for _, i := range res.Inputs {
					fmt.Printf("   %s:\n", (i.Service))
					for _, s := range i.Streams {
						fmt.Printf("      - %s\n", (s))
					}
				}
			}

			if len(res.Outputs) > 0 {
				fmt.Printf("%s\n", style.Gray.Render("\nOutputs:"))
				for _, o := range res.Outputs {
					fmt.Printf("   - %s\n", (o))
				}
			}

			return nil
		},
	}
	infoCmd.Flags().IntVarP(&roverIndex, "rover", "r", 0, "The index of the rover to upload to, between 1 and 20")
	infoCmd.Flags().StringVarP(&roverdHost, "host", "", "", "The roverd endpoint to connect to (if not using --rover)")
	infoCmd.Flags().StringVarP(&roverdUsername, "username", "u", "debix", "The username to use to connect to the roverd endpoint")
	infoCmd.Flags().StringVarP(&roverdPassword, "password", "p", "debix", "The password to use to connect to the roverd endpoint")

	rootCmd.AddCommand(infoCmd)
}
