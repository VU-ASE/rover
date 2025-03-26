package events

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/VU-ASE/rover/roverctl/src/configuration"
	peerconnection "github.com/VU-ASE/rover/roverctl/src/proxy/peerconnection"
	state "github.com/VU-ASE/rover/roverctl/src/proxy/state"

	pb_control "github.com/VU-ASE/rovercom/packages/go/control"
	rtc "github.com/VU-ASE/roverrtc/src"
	"github.com/rs/zerolog/log"

	"github.com/pion/webrtc/v4"
)

// Called when a client sends an offer to the HTTP server
func OnClientSDPReceived(sdp rtc.RequestSDP, state *state.ServerState) ([]byte, error) {
	// Create a new RTCPeerConnection
	rtc, err := peerconnection.CreateFromOffer(sdp.Offer, sdp.Id, state.PeerConfig, state.RtcApi)
	if err != nil {
		return nil, err
	}

	log := rtc.Log()

	// Register connection state change handler
	rtc.Pc.OnConnectionStateChange(onClientConnectionChange(rtc, state))

	// Add rtc to list of client connections
	err = state.ConnectedPeers.Add(sdp.Id, rtc, false)
	if err != nil {
		return nil, err
	}

	log.Debug().Msg("Received SDP offer from client")

	// Register data channel creation and other handlers
	OnClientSDPReturned(rtc, state)

	// Send answer back to client
	payload, err := json.Marshal(rtc.Pc.LocalDescription())
	if err != nil {
		return nil, err
	}

	return payload, nil
}

// Called when a client sends an ICE candidate to the HTTP server
func OnClientICEReceived(ice rtc.RequestICE, state *state.ServerState) ([]byte, error) {
	// Get connection from list of connections
	rtc := state.ConnectedPeers.Get(ice.Id)
	if rtc == nil {
		return nil, fmt.Errorf("Client connection with id %s does not exist", ice.Id)
	}

	log := rtc.Log()

	// Add to list of remote candidates
	if err := rtc.Pc.AddICECandidate(ice.Candidate); err != nil {
		log.Err(err).Msg("Could not add ICE candidate to client connection")
		return nil, err
	}

	// Return all candidates to client
	payload, err := json.Marshal(rtc.GetAllLocalCandidates())
	if err != nil {
		return nil, err
	}

	return payload, nil
}

// Create handlers for data channels
func OnClientSDPReturned(r *rtc.RTC, state *state.ServerState) {
	log := r.Log()

	// Register data channel creation
	r.Pc.OnDataChannel(func(d *webrtc.DataChannel) {
		log.Debug().Str("label", d.Label()).Msg("Client data channel was created")

		// Register channel opening handling
		d.OnOpen(func() {
			log.Debug().Str("label", d.Label()).Msg("Client channel was opened for communication")

			switch d.Label() {
			case configuration.PROXY_CONTROL_CHAN_LABEL:
				registerClientControlMessage(r, d, state)
			case configuration.PROXY_DATA_CHAN_LABEL:
				registerClientDataMessage(r, d, state)
			default:
				log.Warn().Str("label", d.Label()).Msg("Unknown client channel was opened for communication")
			}
		})
	})
}

// Used to bring newly connected clients up to speed with the current state of the server
func onClientConnectionChange(client *rtc.RTC, state *state.ServerState) func(webrtc.PeerConnectionState) {
	log := client.Log()

	return func(s webrtc.PeerConnectionState) {
		log.Debug().Str("newState", s.String()).Msg("Client connection changed to new state")

		if s == webrtc.PeerConnectionStateDisconnected || s == webrtc.PeerConnectionStateClosed || s == webrtc.PeerConnectionStateFailed {
			// Remove the client from the list of connected clients
			_ = state.ConnectedPeers.Remove(client.Id)
			client.Destroy()
		} else if s == webrtc.PeerConnectionStateConnected {
			// Check if there already is a car connected, and send its car state if so
			state.Lock.RLock()
			car := state.ConnectedPeers.Get(configuration.PROXY_CAR_ID)
			state.Lock.RUnlock()
			if car == nil {
				return
			}

			// Create proto message to notify client that a car is connected before they were connected
			notification := pb_control.ConnectionState{
				Client:          car.Id,
				Connected:       car.Pc.ConnectionState() == webrtc.PeerConnectionStateConnected,
				TimestampOffset: car.TimestampOffset,
			}

			// Sleep for 2 seconds to let the webcontroller set up the correct data channel handlers to process our message
			time.Sleep(2 * time.Second)
			log.Debug().Msg("Notifying client of connected car")

			// Notify the client of the car state
			err := client.SendControlData(&notification)
			if err != nil {
				log.Err(err).Msg("Could not notify connected client of connected car")
			}
		}
	}
}

//
// Register data channel message handlers
//

func registerClientControlMessage(client *rtc.RTC, dc *webrtc.DataChannel, state *state.ServerState) {
	client.ControlChannel = dc

	// Register text message handling
	dc.OnMessage(func(msg webrtc.DataChannelMessage) {
		// implement the control protocol here
		log.Warn().Msg("Client should not send control messages, yet. Dropping.")
	})
}

func registerClientDataMessage(client *rtc.RTC, dc *webrtc.DataChannel, state *state.ServerState) {
	client.DataChannel = dc

	// Register text message handling
	dc.OnMessage(func(msg webrtc.DataChannelMessage) {
		log.Debug().Int("length", len(msg.Data)).Msg("Forwarding data: client --> rover")

		// Thank you for the message, we just forward it to the car who can parse it and act accordingly
		car := state.ConnectedPeers.Get(configuration.PROXY_CAR_ID)
		if car == nil {
			// No car to forward to, let the client know
			notification := pb_control.ControlError{
				Message: "No car connected to forward data to",
			}
			log.Warn().Msg("No car connected to forward data to")
			// Best-effort, we don't care if the message fails to send
			_ = client.SendControlData(&notification)
			return
		}

		// Forward it baby
		err := car.SendDataBytes(msg.Data)
		if err != nil {
			log.Err(err).Msg("Could not forward data to car")
		}
	})
}
