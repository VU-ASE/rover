package utils

import (
	"encoding/json"
	"fmt"
)

// Define the structure of the JSON input
type ValidationError struct {
	ValidationErrors struct {
		UnmetStreams []struct {
			Source string `json:"source"`
			Target string `json:"target"`
			Stream string `json:"stream"`
		} `json:"unmet_streams"`
	} `json:"validation_errors"`
}

func TransformValidationError(inputJSON string) ([]string, error) {
	var data ValidationError

	// Parse the JSON input
	err := json.Unmarshal([]byte(inputJSON), &data)
	if err != nil {
		return nil, err
	}

	// Transform unmet streams into nice strings
	var result []string
	for _, stream := range data.ValidationErrors.UnmetStreams {
		description := fmt.Sprintf("Service '%s' depends on input stream '%s' from service '%s', but service '%s' was not enabled",
			stream.Source, stream.Stream, stream.Target, stream.Target)
		result = append(result, description)
	}

	return result, nil
}
