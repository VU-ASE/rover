package events

import (
	"encoding/json"
	"fmt"

	"github.com/VU-ASE/rover/roverctl/src/configuration"
	peerconnection "github.com/VU-ASE/rover/roverctl/src/proxy/peerconnection"
	state "github.com/VU-ASE/rover/roverctl/src/proxy/state"
	pb_control "github.com/VU-ASE/rovercom/packages/go/control"
	rtc "github.com/VU-ASE/roverrtc/src"
	"github.com/pion/webrtc/v4"
	"github.com/rs/zerolog/log"
)

// Called when a car sends an offer to the HTTP server
func OnCarSDPReceived(sdp rtc.RequestSDP, receivedAt int64, state *state.ServerState) ([]byte, error) {
	// Create a new RTCPeerConnection
	rtc, err := peerconnection.CreateFromOffer(sdp.Offer, sdp.Id, state.PeerConfig, state.RtcApi)
	if err != nil {
		return nil, err
	}

	// Set the timestamp offset based on the registration time and the time this offer was received
	rtc.TimestampOffset = sdp.Timestamp - receivedAt

	log := rtc.Log()

	// Register event handlers from now on
	rtc.Pc.OnConnectionStateChange(onCarConnectionChange(rtc, state))

	log.Debug().Msg("Received SDP offer from car")

	// Add rtc to list of car connections (there can be only one car connection)
	err = state.ConnectedPeers.Add(configuration.PROXY_CAR_ID, rtc, true)
	if err != nil {
		return nil, err
	}

	// Register data channel creation and other handlers
	OnCarSDPReturned(rtc, state)

	// Send answer back to car
	payload, err := json.Marshal(rtc.Pc.LocalDescription())
	if err != nil {
		return nil, err
	}

	return payload, nil
}

// Called when a car sends an ICE candidate to the HTTP server
func OnCarICEReceived(ice rtc.RequestICE, state *state.ServerState) ([]byte, error) {
	// Get connection from list of connections
	rtc := state.ConnectedPeers.Get(configuration.PROXY_CAR_ID)
	if rtc == nil {
		return nil, fmt.Errorf("Car connection with id %s does not exist", ice.Id)
	}

	log := rtc.Log()

	// Add to list of remote candidates
	if err := rtc.Pc.AddICECandidate(ice.Candidate); err != nil {
		return nil, err
	}

	// Return all candidates to client
	payload, err := json.Marshal(rtc.GetAllLocalCandidates())
	log.Debug().Msg("Got all local candidates")

	if err != nil {
		return nil, err
	}
	return payload, nil
}

// Register data channel creation and other handlers
func OnCarSDPReturned(r *rtc.RTC, state *state.ServerState) {
	log := r.Log()

	// Register data channel creation
	r.Pc.OnDataChannel(func(d *webrtc.DataChannel) {
		log.Debug().Str("label", d.Label()).Msg("Car channel was created")

		// Register channel opening handling
		d.OnOpen(func() {
			log.Debug().Str("label", d.Label()).Msg("Car channel was opened for communication")

			switch d.Label() {
			case configuration.PROXY_CONTROL_CHAN_LABEL:
				registerCarControlMessage(r, d)
			case configuration.PROXY_DATA_CHAN_LABEL:
				registerCarDataMessage(r, d, state)
			default:
				log.Warn().Str("label", d.Label()).Msg("Unknown car channel was opened for communication")
			}
		})
	})
}

// Inform clients when a car connects/disconnects
func onCarConnectionChange(car *rtc.RTC, state *state.ServerState) func(webrtc.PeerConnectionState) {
	log := car.Log()

	return func(s webrtc.PeerConnectionState) {
		log.Debug().Msgf("Car connection changed to new state %s", s.String())

		notification := pb_control.ConnectionState{
			Connected:       s == webrtc.PeerConnectionStateConnected,
			Client:          configuration.PROXY_CAR_ID,
			TimestampOffset: car.TimestampOffset,
		}

		// Notify all clients of the new connection state
		state.ConnectedPeers.ForEach(func(id string, r *rtc.RTC) {
			if r.Id != car.Id {
				err := r.SendControlData(&notification)
				if err != nil {
					log.Err(err).Str("clientId", id).Msg("Could not notify connected client of car connection state")
				}
			}
		})

		if s == webrtc.PeerConnectionStateConnected {
			//
			// ...
			// Actions on car connection
			// can go here
			// ...
			//
		} else if s == webrtc.PeerConnectionStateDisconnected || s == webrtc.PeerConnectionStateClosed || s == webrtc.PeerConnectionStateFailed {
			// Car disconnected, remove from list of connected peers
			_ = state.ConnectedPeers.Remove(car.Id)
			car.Destroy()
		}
	}
}

//
// Events based on data channel messages
//

func registerCarControlMessage(car *rtc.RTC, dc *webrtc.DataChannel) {
	car.ControlChannel = dc

	// Register text message handling
	dc.OnMessage(func(msg webrtc.DataChannelMessage) {
		// A car should not send control messages, so we do nothing
		log.Warn().Msg("Received control message, but a car should not send control messages. Dropping.")

		// Notify the car that it should not send control messages, best effort
		notification := pb_control.ControlError{
			Message: "You should not send control messages at this time",
		}
		_ = car.SendControlData(&notification)
	})
}

func registerCarDataMessage(car *rtc.RTC, dc *webrtc.DataChannel, state *state.ServerState) {
	car.DataChannel = dc

	// Register text message handling
	dc.OnMessage(func(msg webrtc.DataChannelMessage) {
		log.Debug().Int("length", len(msg.Data)).Msg("Forwarding data: rover --> client")

		// Forward the message to all clients
		state.ConnectedPeers.ForEach(func(id string, r *rtc.RTC) {
			if id == configuration.PROXY_CAR_ID {
				return
			}

			err := r.SendDataBytes(msg.Data)
			if err != nil {
				log.Err(err).Str("clientId", id).Msg("Could not forward data to client")
			}
		})
	})
}
