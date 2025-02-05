package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type UpdateAvailableAsset struct {
	Name        string `json:"name"`
	Url         string `json:"browser_download_url"`
	ContentType string `json:"content_type"`
	Size        int    `json:"size"`
}

type UpdateAvailable struct {
	Name           string
	Author         string
	Available      bool
	CurrentVersion string
	LatestVersion  string
	Assets         []UpdateAvailableAsset // download URLs for the release assets
}

// CheckForGithubUpdate checks if an update is available for the given repository by checking releases on GitHub
func CheckForGithubUpdate(name string, org string, currentVersion string) (*UpdateAvailable, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/latest", org, name)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch latest release from GitHub: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected response from GitHub: %s", resp.Status)
	}

	var release struct {
		TagName string                 `json:"tag_name"`
		Assets  []UpdateAvailableAsset `json:"assets"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return nil, fmt.Errorf("failed to decode GitHub response: %w", err)
	}

	latestVersion := strings.TrimPrefix(release.TagName, "v")
	currentVersion = strings.TrimPrefix(currentVersion, "v")
	available := latestVersion != currentVersion

	return &UpdateAvailable{
		Name:           name,
		Available:      available,
		CurrentVersion: currentVersion,
		LatestVersion:  latestVersion,
		Assets:         release.Assets,
		Author:         org,
	}, nil
}
