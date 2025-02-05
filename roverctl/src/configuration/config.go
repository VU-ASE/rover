package configuration

import (
	"os"

	"gopkg.in/yaml.v3"
)

//
// Get, create and set all the available configuration options (unrelated to connections)
//

// The file name in the configuration directory where the config is stored
var configFilename = LocalConfigDir() + "/config.yaml"

// All the configuration options that can be accessed throughout the application
type RoverctlConfig struct {
	Author string
	Debug  bool
}

// To read config from disk
func ReadConfig() (RoverctlConfig, error) {
	config := RoverctlConfig{
		Author: "",
		Debug:  false,
	}

	// Check if the file exists
	if _, err := os.Stat(configFilename); os.IsNotExist(err) {
		// If the file does not exist, return an empty array
		return config, nil
	}

	// Read the file
	content, err := os.ReadFile(configFilename)
	if err != nil {
		return config, err
	}

	// Parse the YAML content
	err = yaml.Unmarshal(content, &config)
	return config, err
}

// Marshal the config to YAML and save it to disk
func (c RoverctlConfig) Save() error {
	content, err := yaml.Marshal(c)
	if err != nil {
		return err
	}

	// Write the file, overwriting the existing one
	return os.WriteFile(configFilename, content, 0644)
}
