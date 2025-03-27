package view_incompatible

import (
	"context"
	"fmt"

	"github.com/VU-ASE/rover/roverctl/src/configuration"
	"github.com/VU-ASE/rover/roverctl/src/style"
	"github.com/VU-ASE/rover/roverctl/src/utils"
	view_info "github.com/VU-ASE/rover/roverctl/src/views/info"
)

// Best effort try to warn that roverctl and roverd are incompatible
func WarnOnIncompatible(conn configuration.RoverConnection) {
	// Try to query info
	info := conn.ToApiClient().HealthAPI.StatusGet(
		context.Background(),
	)
	res, _, err := info.Execute()
	if err != nil {
		return
	}

	// Check if the versions are compatible
	if !utils.VersionsEqual(res.Version, view_info.Version) {
		s := style.Warning.Bold(true).Render("WARNING!") + " " + style.Warning.Render("roverctl ("+utils.Version(view_info.Version)+") and roverd ("+utils.Version(res.Version)+") are incompatible") + "\n"
		s += "         -> run " + style.Primary.Render("roverctl update") + " to resolve\n\n"

		fmt.Print(s)
	}
}
