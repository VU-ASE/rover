package utils

import "strings"

// Shorthand to create a version string that starts with "v" so that versions can always be compared string-wise
// and in a way that is compatible with the semver package
func Version(s string) string {
	return strings.ToLower("v" + strings.TrimPrefix(s, "v"))
}

func VersionsEqual(a, b string) bool {
	return Version(a) == Version(b)
}
