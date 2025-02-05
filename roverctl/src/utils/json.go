package utils

import (
	"bytes"
	"encoding/json"
)

// Prints JSON bytes nicely indented, or returns the original bytes as string if it fails
func PrettyJSON(j []byte) string {
	// Create a buffer to hold the pretty-printed JSON
	var prettyJSON bytes.Buffer

	// Use json.Indent to format the JSON
	err := json.Indent(&prettyJSON, j, "", "  ")
	if err != nil {
		return string(j)
	}

	// Return the pretty-printed JSON as a string
	return prettyJSON.String()
}
