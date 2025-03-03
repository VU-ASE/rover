package httpserver

//
// The HTTP server is an entry point for the answer process for both the car and connected clients.
// it is necessary to set up webRTC connections (the signaling process) between clients <-> server and server <-> car.
// several HTTP endpoints are defined to handle the signaling process.
//

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/rs/zerolog/log"

	"net/http"
	"vu/ase/streamserver/src/events"
	"vu/ase/streamserver/src/state"

	rtc "github.com/VU-ASE/roverrtc/src"
)

type EndpointError struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

// Explains usage of an HTTP endpoint. Returns true if the request was a GET request.
func explainPostUsage(w http.ResponseWriter, r *http.Request, usage string) bool {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Header().Set("Content-Type", "application/text")
		_, _ = w.Write([]byte(fmt.Sprintf("Endpoint OK. \n\nTo use this endpoint, send a POST request\n%s", usage)))
		return true
	}
	return false
}

// Template function for creating an HTTP endpoint with error handling and CORS headers
func JSONEndpoint(usage string, handler func(w http.ResponseWriter, r *http.Request) ([]byte, error)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")

		if explainPostUsage(w, r, usage) {
			return
		}

		if payload, err := handler(w, r); err != nil {
			// Log the error to the console and send a HTTP response
			log.Err(err).Str("endpoint", r.URL.Path).Str("method", r.Method).Msg("Could not process request")
			w.WriteHeader(http.StatusInternalServerError)

			// Encode error as JSON
			errorObj := EndpointError{
				Error:   true,
				Message: err.Error(),
			}
			payload, err := json.Marshal(errorObj)
			if err != nil {
				log.Err(err).Str("endpoint", r.URL.Path).Str("method", r.Method).Msg("Could not encode error as JSON")
			} else {
				_, _ = w.Write(payload)
			}
		} else {
			// Let the HTTP client know that the request was successful
			w.WriteHeader(http.StatusCreated)
			_, _ = w.Write(payload)
		}
	}
}

// Configure the HTTP server to listen for incoming connections on the configured endpoints
func Serve(serverAddress string, state *state.ServerState) error {

	//
	// Client endpoints
	//

	// To show a welcome message when someone connects to the base URL through the browser
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	// To retrieve an SDP offer (and send back an SDP answer)
	http.HandleFunc("/client/sdp", JSONEndpoint("[💻 CLIENT ONLY]: Send your SDP offer as a JSON object", func(w http.ResponseWriter, r *http.Request) ([]byte, error) {
		// Parse offer from request body
		request := rtc.RequestSDP{}
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			return nil, err
		}

		// Process offer
		return events.OnClientSDPReceived(request, state)
	}))

	// To retrieve an ICE candidate (and send back an ICE candidate)
	http.HandleFunc("/client/ice", JSONEndpoint("[💻 CLIENT ONLY]: Send your ICE candidate as a JSON object", func(w http.ResponseWriter, r *http.Request) ([]byte, error) {
		// Parse ICE from request body
		request := rtc.RequestICE{}
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			return nil, err
		}

		return events.OnClientICEReceived(request, state)
	}))

	//
	// Car endpoints
	//

	// To retrieve an SDP offer (and send back an SDP answer)
	http.HandleFunc("/car/sdp", JSONEndpoint("[🚗 CAR ONLY]: Send your SDP offer as a JSON object", func(w http.ResponseWriter, r *http.Request) ([]byte, error) {
		// Record the timestamp at which this request was received
		receivedAt := time.Now().UnixMilli()

		// Parse offer from request body
		request := rtc.RequestSDP{}
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			return nil, err
		}

		return events.OnCarSDPReceived(request, receivedAt, state)
	}))

	// To retrieve an ICE candidate (and send back an ICE candidate)
	http.HandleFunc("/car/ice", JSONEndpoint("[🚗 CAR ONLY]: Send your ICE candidate as a JSON object", func(w http.ResponseWriter, r *http.Request) ([]byte, error) {
		// Parse ICE from request body
		request := rtc.RequestICE{}
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			return nil, err
		}

		return events.OnCarICEReceived(request, state)
	}))

	// Start HTTP server to accept incoming connections
	log.Info().Msgf("passthrough HTTP listener active on '%s'", serverAddress)

	err := http.ListenAndServe(serverAddress, nil)
	if err != nil {
		return fmt.Errorf("Cannot start HTTP server: %v", err)
	}
	return err
}
