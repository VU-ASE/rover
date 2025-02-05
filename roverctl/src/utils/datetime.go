package utils

import (
	"fmt"
	"time"
)

func FormatDuration(milliseconds int64) string {
	// Convert milliseconds to time.Duration
	duration := time.Duration(milliseconds) * time.Millisecond

	// Extract hours, minutes, seconds
	hours := duration / time.Hour
	duration %= time.Hour
	minutes := duration / time.Minute
	duration %= time.Minute
	seconds := duration / time.Second

	// Build the formatted string
	formatted := ""
	if hours > 0 {
		formatted += fmt.Sprintf("%dh ", hours)
	}
	if minutes > 0 {
		formatted += fmt.Sprintf("%dm ", minutes)
	}
	if seconds > 0 || formatted == "" { // Always include seconds if no other units are present
		formatted += fmt.Sprintf("%ds", seconds)
	}

	return formatted
}
