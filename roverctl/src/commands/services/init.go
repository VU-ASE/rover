package command_services

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/spf13/cobra"

	state "github.com/VU-ASE/rover/roverctl/src/state"
	"github.com/VU-ASE/rover/roverctl/src/style"
	git "github.com/go-git/go-git/v5"
)

var presets = []string{"go", "c", "python", "cpp"}

func addInit(rootCmd *cobra.Command) {
	// General flags
	var name string
	var source string

	// services command
	var infoCmd = &cobra.Command{
		Use:   "init [" + strings.Join(presets, "|") + "]",
		Short: "Create a new service in your current working directory, based on a template",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return fmt.Errorf("Specify a preset to use, one of: %v", strings.Join(presets, ", "))
			}
			// Check if the preset is valid
			for _, preset := range presets {
				if args[0] == preset {
					return nil
				}
			}
			return fmt.Errorf("Invalid preset, must be one of: %v", strings.Join(presets, ", "))
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			s := state.Get()
			if s == nil || s.Config.Author == "" {
				s := "Uh oh, " + style.Error.Render("roverctl does not know who you are yet") + "!\n"
				s += "To get started, just run " + style.Primary.Render("roverctl author --set <NAME>") + " and try again.\n"
				fmt.Print(s)
				return nil
			}

			//
			// Validation
			//

			if len(name) < 3 {
				fmt.Println("Cannot initialize service: service names must be at least 3 characters long")
				return nil
			}

			// Check if a folder with this name already exists in the current working directory
			_, err := os.Stat(name)
			if err == nil {
				fmt.Printf("Cannot initialize service: a folder with the name '%s' already exists in the current directory\n", name)
				return nil
			}

			// Can only contain lowercase letters and hyphens
			valid := regexp.MustCompile(`^[a-z0-9-]*$`).MatchString(name)
			if !valid {
				fmt.Println("Cannot initialize service: service names can only contain lowercase letters and hyphens")
				return nil
			}

			if source == "" {
				fmt.Printf("Cannot initialize service: source URL is invalid")
				return nil
			}
			if strings.Contains(source, "username") || strings.Contains(source, "repository") {
				fmt.Printf("Cannot initialize service: replace 'username' and 'repository' with your actual GitHub username/organization name and repository name")
				return nil
			}
			if strings.Contains(source, "https://") || strings.Contains(source, "http://") || strings.Contains(source, "www.") {
				fmt.Printf("Cannot initialize service: do not include the protocol or 'www.' in the URL")
				return nil
			}

			preset := args[0]
			err = initializeTemplate(preset, name, s.Config.Author, "0.0.1", source)
			if err != nil {
				fmt.Printf("Cannot not initialize service: %s\n", err.Error())
			} else {
				fmt.Printf("Service %s was initialized successfully\n", style.Primary.Render(name))
				fmt.Printf("Enter %s in your terminal and fire up the devcontainer to get started!\n", style.Gray.Render("cd "+name))
			}

			return nil
		},
	}
	infoCmd.Flags().StringVarP(&name, "name", "n", "", "The name of the service to create")
	infoCmd.Flags().StringVarP(&source, "source", "s", "", "The source (github URL) of the service to create")
	infoCmd.MarkFlagRequired("name")
	infoCmd.MarkFlagRequired("source")

	rootCmd.AddCommand(infoCmd)
}

// Downloads a selected template from a repository and places it in the destination folder
func downloadTemplate(repository string, destination string) error {
	// Temp directory to clone in
	tmp, err := os.MkdirTemp("", "roverctl-service-init")
	if err != nil {
		return fmt.Errorf("Could not clone template: %v", err)
	}

	// Clone the repository in temp
	_, err = git.PlainClone(tmp, false, &git.CloneOptions{
		URL: repository,
	})
	if err != nil {
		return fmt.Errorf("Could not clone template: %v", err)
	}
	// Remove the .git folder from the template
	err = os.RemoveAll(filepath.Join(tmp, ".git"))
	if err != nil {
		return fmt.Errorf("Could not remove .git folder: %v", err)
	}

	// Move the template to the destination. Using linux 'mv' command because Go's os.Rename requires
	// src and dest to be part of the same mount point, which in some cases is not always true.
	// Testing showed that using os.Rename on some systems would yield "invalid cross-device link" error.
	cmd := exec.Command("mv", tmp, destination)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("Could not move template: %v", err)
	}

	return err
}

func initializeTemplate(preset string, name string, author string, version string, source string) error {
	// Based on the programming language chosen, download a specific template and replace the magic strings in it
	templateRepo := ""
	switch preset {
	case presets[0]:
		templateRepo = "https://github.com/VU-ASE/service-template-go"
	case presets[1]:
		templateRepo = "https://github.com/VU-ASE/service-template-c"
	case presets[2]:
		templateRepo = "https://github.com/VU-ASE/service-template-python"
	case presets[3]:
		templateRepo = "https://github.com/VU-ASE/service-template-cpp"
	}

	if templateRepo == "" {
		return fmt.Errorf("No template found for preset %s", preset)
	}

	fmt.Printf("Downloading service template from %s\n", templateRepo)
	// Get the current working directory
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("Could not get current working directory: %v", err)
	}

	dest := filepath.Join(cwd, name)
	err = downloadTemplate(templateRepo, dest)
	if err != nil {
		return err
	}

	// Strings to be replaced
	toReplace := map[string]string{
		"SERVICE_NAME":    name,
		"SERVICE_AUTHOR":  author,
		"SERVICE_VERSION": version,
		"SERVICE_SOURCE":  source,
	}

	for key, value := range toReplace {
		// Escape key and value
		escapedKey := escapeShellString(key)
		escapedValue := escapeShellString(value)

		// Build the `find` and `sed` command
		// This replaces `SERVICE_NAME` with `MyService` in all files
		cmd := exec.Command("bash", "-c", fmt.Sprintf(
			`find . -type f -exec sed -i.bak 's/%s/%s/g' {} + && find . -type f -name "*.bak" -delete`,
			escapedKey, escapedValue))

		// Set the current working directory (optional, defaults to where the tool is run)
		cmd.Dir = dest

		if err := cmd.Run(); err != nil {
			return fmt.Errorf("Error running replacement command: %v\n", err)
		}
	}

	return nil
}

func escapeShellString(input string) string {
	// Escape special characters for sed
	replacer := strings.NewReplacer(
		`&`, `\&`, // Escape &
		`/`, `\/`, // Escape /
		`'`, `'\''`, // Escape single quotes
	)
	return replacer.Replace(input)
}
