package utils

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/VU-ASE/rover/roverctl/src/openapi"
	"gopkg.in/yaml.v3"
)

type ServiceFqn struct {
	Author  string
	Name    string
	Version string
}

func ServiceFqnToStr(s ServiceFqn) string {
	return fmt.Sprintf("%s/%s@%s", s.Author, s.Name, s.Version)
}

// Parsed from the service YAML files (simplified)
type ServiceInformation struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
}

func GetServiceInformation(path string) (*ServiceInformation, error) {
	// Try to parse the yaml for this path
	yamlPath := filepath.Join(path, "service.yaml")
	content, err := os.ReadFile(yamlPath)
	if err != nil {
		return nil, err
	}

	// Parse the yaml
	info := ServiceInformation{}
	err = yaml.Unmarshal(content, &info)
	if err != nil || info.Name == "" || info.Version == "" {
		return nil, fmt.Errorf("Service YAML was invalid")
	}

	return &info, nil
}

func FqnsEqual(a openapi.FullyQualifiedService, b openapi.FullyQualifiedService) bool {
	return a.Name == b.Name && a.Author == b.Author && a.Version == b.Version
}
