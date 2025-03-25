package utils

import (
	"bytes"
	"fmt"
	"net"
	"os/exec"
	"runtime"
	"strings"
)

// getLocalIP finds and returns the first local (192.168.x.x) IP address.
func GetLocalIP() (string, error) {
	// Get all network interfaces
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", fmt.Errorf("failed to get network interfaces: %v", err)
	}

	for _, iface := range interfaces {
		// Ignore interfaces that are down or loopback
		if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback != 0 {
			continue
		}

		// Get addresses associated with the interface
		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}

		for _, addr := range addrs {
			// Parse IP address
			ipNet, ok := addr.(*net.IPNet)
			if !ok || ipNet.IP.IsLoopback() {
				continue
			}

			ip := ipNet.IP.To4()
			if ip == nil {
				continue // Ignore IPv6 addresses
			}

			// Check if it's a 192.168.x.x address
			if strings.HasPrefix(ip.String(), "192.168.") {
				return ip.String(), nil
			}
		}
	}

	return "", fmt.Errorf("no local (192.168.x.x) IP address found")
}

// IsPortAvailable checks if a port is available on the local machine
func IsPortAvailable(port int) bool {
	addr := fmt.Sprintf(":%d", port)
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return false // Port is in use or inaccessible
	}
	_ = ln.Close()
	return true
}

func GetProcessUsingPort(port int) (string, error) {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "darwin", "linux":
		// Example: lsof -i :8080
		cmd = exec.Command("lsof", "-i", fmt.Sprintf(":%d", port))
	default:
		return "", fmt.Errorf("unsupported OS: %s", runtime.GOOS)
	}

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("error running lsof: %v - output: %s", err, out.String())
	}

	return out.String(), nil
}
