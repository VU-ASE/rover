package peerconnection

import (
	"fmt"

	rtc "github.com/VU-ASE/roverrtc/src"

	"github.com/pion/webrtc/v4"
)

// Create new RTC connection from an SDP offer
func CreateFromOffer(offer webrtc.SessionDescription, id string, peerConfig webrtc.Configuration, webrtcApi *webrtc.API) (*rtc.RTC, error) {
	// New RTC object that might be returned
	rtc := rtc.NewRTC(id)

	log := rtc.Log()

	// Create a local PeerConnection
	peerConnection, err := webrtcApi.NewPeerConnection(peerConfig)
	if err != nil {
		return nil, fmt.Errorf("Could not create peer connection: %v", err)
	}
	rtc.Pc = peerConnection

	// Fetch all pending ICE candidates (we will wait for the ICE gathering to complete before sending the answer)
	rtc.Pc.OnICECandidate(func(c *webrtc.ICECandidate) {
		if c != nil {
			rtc.AddLocalCandidate(c.ToJSON())
		}
	})

	// Set the remote SessionDescription based on the incoming offer
	err = rtc.Pc.SetRemoteDescription(offer)
	if err != nil {
		return nil, fmt.Errorf("Could not set remote description: %v", err)
	}

	// Create answer for the remote peer (to confirm the connection)
	answer, err := rtc.Pc.CreateAnswer(nil)
	if err != nil {
		return nil, fmt.Errorf("Could not create answer: %v", err)
	}

	// Create channel that is blocked until ICE Gathering is complete (don't use trickling ICE)
	gatherComplete := webrtc.GatheringCompletePromise(rtc.Pc)

	// Sets the LocalDescription, and starts our UDP listeners
	err = rtc.Pc.SetLocalDescription(answer)
	if err != nil {
		return nil, fmt.Errorf("Could not set local description: %v", err)
	}

	// Block until ICE Gathering is complete, disabling trickle ICE so that we can send the answer as one blob
	<-gatherComplete
	// from this point on, the ICE candidates are complete (and we don't need locks anymore)
	log.Debug().Msg("ICE gathering completed")

	return rtc, nil
}
