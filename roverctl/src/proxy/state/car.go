package state

import (
	"fmt"
	"sync"

	rtc "github.com/VU-ASE/roverrtc/src"
	"github.com/pion/ice/v3"
	"github.com/pion/webrtc/v4"
	"github.com/rs/zerolog/log"
)

// This makes the server easier to test and mock and also allows us to
// add more fields to the server state in the future.
type ServerState struct {
	RtcApi         *webrtc.API
	ConnectedPeers *rtc.RTCMap
	ServerRealIp   string               // the real IP of the server (the host machine, when running through Docker)
	Lock           *sync.RWMutex        // to make sure ICE candidates can be managed concurrently
	PeerConfig     webrtc.Configuration // reusable peer configuration that describes which STUN servers to use (if using WAN)
}

// Create a new proxy server, specifying the IP address to bind to and the port to use for muxing UDP packets
func New(ip string, muxUdpPort int, useWan bool) (*ServerState, error) {
	s := webrtc.SettingEngine{}

	// Need to know our own IP address, so that we can fetch ICE candidates correctly
	if ip == "" {
		return nil, fmt.Errorf("Could not start proxy server, no IP address was specified")
	}
	if muxUdpPort <= 0 || muxUdpPort > 56000 {
		return nil, fmt.Errorf("Could not start proxy server, invalid port number: %d", muxUdpPort)
	}

	s.SetNAT1To1IPs([]string{ip}, webrtc.ICECandidateTypeHost)
	mux, err := ice.NewMultiUDPMuxFromPort(muxUdpPort)
	if err != nil {
		return nil, fmt.Errorf("Could not create UDP mux: %v", err)
	}
	s.SetICEUDPMux(mux)
	log.Debug().Msgf("webRTC udp listener active on '%s:%d'", ip, muxUdpPort)

	// Create a local PeerConnection
	api := webrtc.NewAPI(webrtc.WithSettingEngine(s))

	peerconfig := webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{},
	}
	if useWan {
		peerconfig.ICEServers = append(peerconfig.ICEServers, webrtc.ICEServer{
			URLs: []string{"stun:stun.l.google.com:19302"},
		})
	}

	return &ServerState{
		RtcApi:         api,
		ConnectedPeers: rtc.NewRTCMap(),
		ServerRealIp:   ip,
		Lock:           &sync.RWMutex{},
		PeerConfig:     peerconfig,
	}, nil
}

// Destroy all connections and the server state
func (s *ServerState) Destroy() {
	s.Lock.Lock()
	defer s.Lock.Unlock()

	for _, peer := range s.ConnectedPeers.UnsafeGetAll() {
		_ = s.ConnectedPeers.Remove(peer.Id)
		peer.Destroy()
	}

	log.Debug().Msg("Destroyed server state")
}
