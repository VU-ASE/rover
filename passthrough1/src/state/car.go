package state

import (
	"fmt"
	"os"
	"sync"
	livestreamconfig "vu/ase/streamserver/src/config"

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

func NewServerState(useWan bool) (*ServerState, error) {
	s := webrtc.SettingEngine{}

	// This is necessary when running through Docker, so that the ICE candidates can be resolved
	serverIp := os.Getenv("ASE_SERVER_IP")
	if serverIp == "" {
		return nil, fmt.Errorf("ASE_SERVER_IP environment variable not set. Please set it to your local IP address (192.168.0.XXX)")
	} else {
		log.Info().Msgf("passthrough webRTC listener active on '%s:%d'", serverIp, livestreamconfig.MuxUdpPort)
	}
	s.SetNAT1To1IPs([]string{serverIp}, webrtc.ICECandidateTypeHost)

	mux, err := ice.NewMultiUDPMuxFromPort(livestreamconfig.MuxUdpPort)
	if err != nil {
		return nil, fmt.Errorf("Could not create UDP mux: %v", err)
	}
	s.SetICEUDPMux(mux)

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
		ServerRealIp:   serverIp,
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

	log.Info().Msg("Destroyed server state")
}
