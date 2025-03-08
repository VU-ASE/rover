package proxy

import (
	"fmt"
	"os"
	"strings"

	httpserver "github.com/VU-ASE/rover/roverctl/src/proxy/httpserver"
	state "github.com/VU-ASE/rover/roverctl/src/proxy/state"

	rtc "github.com/VU-ASE/roverrtc/src"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Run(ip string, muxUdpPort int, useWan bool, verbose bool) error {
	setupLogging(verbose)

	state, err := state.New(ip, muxUdpPort, useWan)
	if err != nil {
		return fmt.Errorf("Could not create proxy server state: %v", err)
	}

	// Create a map to hold all active connections
	connectedPeers := rtc.NewRTCMap()
	// Clean up connections when the server is shut down
	defer state.Destroy()

	// Global server state
	state.ConnectedPeers = connectedPeers

	// Always serve HTTP on port 7500
	httpAddr := fmt.Sprintf("%s:%d", "0.0.0.0", 7500)
	return httpserver.Serve(httpAddr, state)
}

// Configures log level and output
func setupLogging(verbose bool) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs

	// Set up custom caller prefix
	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		path := strings.Split(file, "/")
		// only take the last three elements of the path
		filepath := strings.Join(path[len(path)-3:], "/")
		return fmt.Sprintf("[%s] %s:%d", "proxy", filepath, line)
	}
	outputWriter := zerolog.ConsoleWriter{Out: os.Stderr}
	log.Logger = log.Output(outputWriter).With().Caller().Logger()

	// Set log level
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if verbose {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Debug().Msg("Debug logs enabled")
	}
}
