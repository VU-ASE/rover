package utils

import (
	"fmt"
	"net"
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
