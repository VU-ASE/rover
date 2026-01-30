package utils

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"regexp"
	"time"
	"runtime"
)

// A regular expression to match IPv4 addresses in ping output
var theIPv4Regex = regexp.MustCompile(`$begin:math:text$\(\\d\{1\,3\}\(\?\:\\\.\\d\{1\,3\}\)\{3\}\)$end:math:text$`)

func ResolveHostWithPing(host string) (string, error) {
	// We use a context with timeout to avoid hanging indefinitely on the ping command
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// We prepare the ping command based on what OS is currently running roverctl
	args := pingArgs(host)
	cmd := exec.CommandContext(ctx, "ping", args...)

	// Capture the output
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	// Run the ping command, but do not return an error yet
	// Since ping might fail even if mDNS resolution worked
	// We will check the output for an IP address, and if we cannot find one, we return an error
	if err := cmd.Run(); err != nil {
		_ = err
	}

	// Search the output for an IPv4 address
	matches := theIPv4Regex.FindStringSubmatch(out.String())
	if len(matches) >= 2 {
		return matches[1], nil
	}

	// If we reach here, we could not find an IP address in the ping output
	return "", fmt.Errorf("could not resolve host %s via ping, output: %s", host, out.String())
}

func pingArgs(host string) []string {
	// Different OSes have different arguments for ping
	switch runtime.GOOS {
		// MacOS
		case "darwin":
			return []string{"-c", "1", "-W", "1000", host}
		// Linux
		default:
			return []string{"-c", "1", "-w", "1", host}
	}
}



