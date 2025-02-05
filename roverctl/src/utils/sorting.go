package utils

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// Parse a semantic version string into three integers: major, minor, patch
func parseVersion(version string) (int, int, int, error) {
	parts := strings.Split(version, ".")
	if len(parts) < 3 {
		return 0, 0, 0, fmt.Errorf("invalid version: %s", version)
	}

	major, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, 0, fmt.Errorf("invalid major version: %s", parts[0])
	}

	minor, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, 0, 0, fmt.Errorf("invalid minor version: %s", parts[1])
	}

	patch, err := strconv.Atoi(parts[2])
	if err != nil {
		return 0, 0, 0, fmt.Errorf("invalid patch version: %s", parts[2])
	}

	return major, minor, patch, nil
}

// Compare two versions
func compareVersions(v1, v2 string) bool {
	major1, minor1, patch1, _ := parseVersion(v1)
	major2, minor2, patch2, _ := parseVersion(v2)

	// Compare versions
	if major1 != major2 {
		return major1 > major2
	}
	if minor1 != minor2 {
		return minor1 > minor2
	}
	return patch1 > patch2
}

// SortByVersion sorts a slice of strings by semantic version number, highest first
func SortByVersion(unsorted []string) []string {
	sorted := make([]string, len(unsorted))
	copy(sorted, unsorted)

	sort.Slice(sorted, func(i, j int) bool {
		return compareVersions(sorted[i], sorted[j])
	})

	return sorted
}
