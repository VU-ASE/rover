package utils

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"syscall"
)

// ExecuteShellCommand replaces the current process with a shell that runs the given command.
// This supports pipes, redirection, and all shell features.
func ExecuteShellCommand(command string) error {
	var shell, flag string

	// Determine the shell and its command flag based on the operating system
	if runtime.GOOS == "windows" {
		shell = "cmd.exe"
		flag = "/c"
	} else {
		shell = "/bin/sh"
		flag = "-c"
	}

	// Replace the current process with the shell and command
	err := syscall.Exec(shell, []string{shell, flag, command}, os.Environ())
	if err != nil {
		return fmt.Errorf("failed to execute shell command: %w", err)
	}

	return nil
}

func OpenBrowser(url string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "linux":
		cmd = exec.Command("xdg-open", url)
	case "darwin":
		cmd = exec.Command("open", url)
	default:
		return fmt.Errorf("unsupported platform")
	}

	return cmd.Start()
}
