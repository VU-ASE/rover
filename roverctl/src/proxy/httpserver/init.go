package httpserver

//
// The HTTP server is an entry point for the answer process for both the car and connected clients.
// it is necessary to set up webRTC connections (the signaling process) between clients <-> server and server <-> car.
// several HTTP endpoints are defined to handle the signaling process.
//

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	events "github.com/VU-ASE/rover/roverctl/src/proxy/events"
	state "github.com/VU-ASE/rover/roverctl/src/proxy/state"
	rtc "github.com/VU-ASE/roverrtc/src"
	"github.com/rs/zerolog/log"
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

		log.Debug().Msgf("HTTP endpoint %s called", r.URL.Path)

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
func Serve(serverAddress string, s *state.ServerState) error {

	//
	// Client endpoints
	//

	// To show a welcome message when someone connects to the base URL through the browser
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
		<!doctype html>
<html lang="en">

<!--
    Credits for this template: https://github.com/leemunroe/responsive-html-email-template/blob/master/email.html
-->

<head>
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
    <title>ASE debugging proxy</title>
    <style media="all" type="text/css">
        /* -------------------------------------
    GLOBAL RESETS
------------------------------------- */

        body {
            font-family: Helvetica, sans-serif;
            -webkit-font-smoothing: antialiased;
            font-size: 16px;
            line-height: 1.3;
            -ms-text-size-adjust: 100%;
            -webkit-text-size-adjust: 100%;
        }

        table {
            border-collapse: separate;
            mso-table-lspace: 0pt;
            mso-table-rspace: 0pt;
            width: 100%;
        }

        table td {
            font-family: Helvetica, sans-serif;
            font-size: 16px;
            vertical-align: top;
        }

        /* -------------------------------------
    BODY & CONTAINER
------------------------------------- */

        body {
            background-color: #f4f5f6;
            margin: 0;
            padding: 0;
        }

        .body {
            background-color: #f4f5f6;
            width: 100%;
        }

        .container {
            margin: 0 auto !important;
            max-width: 600px;
            padding: 0;
            padding-top: 24px;
            width: 600px;
        }

        .content {
            box-sizing: border-box;
            display: block;
            margin: 0 auto;
            max-width: 600px;
            padding: 0;
        }

        /* -------------------------------------
    HEADER, FOOTER, MAIN
------------------------------------- */

        .main {
            background: #ffffff;
            border: 1px solid #eaebed;
            border-radius: 16px;
            width: 100%;
        }

        .wrapper {
            box-sizing: border-box;
            padding: 24px;
        }

        .footer {
            clear: both;
            padding-top: 24px;
            text-align: center;
            width: 100%;
        }

        .footer td,
        .footer p,
        .footer span,
        .footer a {
            color: #9a9ea6;
            font-size: 16px;
            text-align: center;
        }

        /* -------------------------------------
    TYPOGRAPHY
------------------------------------- */

        p {
            font-family: Helvetica, sans-serif;
            font-size: 16px;
            font-weight: normal;
            margin: 0;
            margin-bottom: 16px;
        }

        a {
            color: #0867ec;
            text-decoration: underline;
        }

        /* -------------------------------------
    BUTTONS
------------------------------------- */

        .btn {
            box-sizing: border-box;
            min-width: 100% !important;
            width: 100%;
        }

        .btn>tbody>tr>td {
            padding-bottom: 16px;
        }

        .btn table {
            width: auto;
        }

        .btn table td {
            background-color: #ffffff;
            border-radius: 4px;
            text-align: center;
        }

        .btn a {
            background-color: #ffffff;
            border: solid 2px #0867ec;
            border-radius: 4px;
            box-sizing: border-box;
            color: #0867ec;
            cursor: pointer;
            display: inline-block;
            font-size: 16px;
            font-weight: bold;
            margin: 0;
            padding: 12px 24px;
            text-decoration: none;
            text-transform: capitalize;
        }

        .btn-primary table td {
            background-color: #0867ec;
        }

        .btn-primary a {
            background-color: #0867ec;
            border-color: #0867ec;
            color: #ffffff;
        }

        @media all {
            .btn-primary table td:hover {
                background-color: #ec0867 !important;
            }

            .btn-primary a:hover {
                background-color: #ec0867 !important;
                border-color: #ec0867 !important;
            }
        }

        /* -------------------------------------
    OTHER STYLES THAT MIGHT BE USEFUL
------------------------------------- */

        .last {
            margin-bottom: 0;
        }

        .first {
            margin-top: 0;
        }

        .align-center {
            text-align: center;
        }

        .align-right {
            text-align: right;
        }

        .align-left {
            text-align: left;
        }

        .text-link {
            color: #0867ec !important;
            text-decoration: underline !important;
        }

        .clear {
            clear: both;
        }

        .mt0 {
            margin-top: 0;
        }

        .mb0 {
            margin-bottom: 0;
        }

        .preheader {
            color: transparent;
            display: none;
            height: 0;
            max-height: 0;
            max-width: 0;
            opacity: 0;
            overflow: hidden;
            mso-hide: all;
            visibility: hidden;
            width: 0;
        }

        .powered-by a {
            text-decoration: none;
        }

        /* -------------------------------------
    RESPONSIVE AND MOBILE FRIENDLY STYLES
------------------------------------- */

        @media only screen and (max-width: 640px) {

            .main p,
            .main td,
            .main span {
                font-size: 16px !important;
            }

            .wrapper {
                padding: 8px !important;
            }

            .content {
                padding: 0 !important;
            }

            .container {
                padding: 0 !important;
                padding-top: 8px !important;
                width: 100% !important;
            }

            .main {
                border-left-width: 0 !important;
                border-radius: 0 !important;
                border-right-width: 0 !important;
            }

            .btn table {
                max-width: 100% !important;
                width: 100% !important;
            }

            .btn a {
                font-size: 16px !important;
                max-width: 100% !important;
                width: 100% !important;
            }
        }

        /* -------------------------------------
    PRESERVE THESE STYLES IN THE HEAD
------------------------------------- */

        @media all {
            .ExternalClass {
                width: 100%;
            }

            .ExternalClass,
            .ExternalClass p,
            .ExternalClass span,
            .ExternalClass font,
            .ExternalClass td,
            .ExternalClass div {
                line-height: 100%;
            }

            .apple-link a {
                color: inherit !important;
                font-family: inherit !important;
                font-size: inherit !important;
                font-weight: inherit !important;
                line-height: inherit !important;
                text-decoration: none !important;
            }

            #MessageViewBody a {
                color: inherit;
                text-decoration: none;
                font-size: inherit;
                font-family: inherit;
                font-weight: inherit;
                line-height: inherit;
            }
        }
    </style>
</head>

<body>
    <table role="presentation" border="0" cellpadding="0" cellspacing="0" class="body">
        <tr>
            <td>&nbsp;</td>
            <td class="container">
                <div class="content">

                    <!-- START CENTERED WHITE CONTAINER -->
                    <span class="preheader">This is preheader text. Some clients will show this text as a
                        preview.</span>
                    <table role="presentation" border="0" cellpadding="0" cellspacing="0" class="main">

                        <!-- START MAIN CONTENT AREA -->
                        <tr>

                            <svg id="Layer_1" data-name="Layer 1" xmlns="http://www.w3.org/2000/svg"
                                viewBox="0 0 4020 1000">
                                <defs>
                                    <style>
                                        .cls-1 {
                                            fill: #1d1d1b;
                                        }

                                        .cls-1,
                                        .cls-2,
                                        .cls-3 {
                                            stroke-width: 0px;
                                        }

                                        .cls-2 {
                                            fill: #0f0f0f;
                                        }

                                        .cls-3 {
                                            fill: #0089cf;
                                        }
                                    </style>
                                </defs>
                                <g>
                                    <path class="cls-2"
                                        d="M2300.48,665.93v-73.42h174.38c2.57,0,4.74-.89,6.52-2.66,1.77-1.77,2.66-3.95,2.66-6.52s-.89-4.74-2.66-6.52c-1.78-1.77-3.95-2.66-6.52-2.66h-91.78c-22.76,0-42.22-8.08-58.37-24.23-16.15-16.15-24.23-35.61-24.23-58.37s8.08-42.22,24.23-58.37c16.15-16.15,35.61-24.23,58.37-24.23h174.38v73.42h-174.38c-2.57,0-4.74.89-6.52,2.66-1.78,1.78-2.66,3.95-2.66,6.52s.89,4.74,2.66,6.52c1.77,1.78,3.95,2.66,6.52,2.66h91.78c22.76,0,42.22,8.08,58.37,24.23,16.15,16.15,24.23,35.61,24.23,58.37s-8.08,42.22-24.23,58.37c-16.15,16.15-35.61,24.23-58.37,24.23h-174.38Z" />
                                    <path class="cls-2"
                                        d="M2851.17,408.95v73.42h-183.56v18.36h183.56v73.43h-183.56v18.36h183.56v73.42h-256.99v-256.99h256.99Z" />
                                    <polygon class="cls-2"
                                        points="2190.35 408.95 1933.36 665.93 2043.5 665.93 2190.35 519.08 2190.35 592.51 2135.28 592.51 2133.73 592.51 2060.31 665.93 2135.28 665.93 2170.44 665.93 2263.77 665.93 2263.77 408.95 2190.35 408.95" />
                                    <path class="cls-2"
                                        d="M1840.58,665.93l256.99-256.99h9.79l-256.99,256.99h-9.79ZM1880.91,665.93l256.99-256.99h-13.61l-256.99,256.99h13.61ZM1916.77,665.93l256.99-256.99h-18.93l-256.99,256.99h18.93Z" />
                                </g>
                                <g>
                                    <polygon class="cls-1"
                                        points="1264.07 663.59 1168.83 407.84 1226.43 407.84 1299.39 605.99 1370.82 407.84 1409.22 407.84 1317.06 663.59 1264.07 663.59" />
                                    <path class="cls-1"
                                        d="M1439.94,564.51c0,29.95,4.61,51.46,13.06,64.51,9.22,13.82,20.74,23.81,36.1,30.72,15.36,6.91,33.79,9.98,54.53,9.98s39.17-3.84,53.76-10.75c13.82-7.68,24.58-17.66,32.26-30.72,7.68-13.06,11.52-33.79,11.52-63.75v-156.68h-46.08v160.52c0,26.11-4.61,43.78-13.06,52.99-9.22,9.22-20.74,13.82-36.1,13.82s-27.65-4.61-37.63-13.82c-9.98-9.22-14.59-28.42-14.59-56.83v-156.68h-52.99v156.68h-.77Z" />
                                    <path class="cls-3"
                                        d="M2004.44,408.6c3.84,0,7.68,0,6.14,5.38,0,3.84-1.54,5.38-7.68,6.91-18.43,3.07-26.11,6.91-53.76,18.43,13.06-3.84,24.58-6.14,30.72-6.14,3.07,0,4.61,0,4.61,3.07,0,3.84-2.3,6.14-5.38,6.91-16.13,5.38-37.63,15.36-57.6,23.81,20.74-6.14,30.72-6.14,35.33-6.14,3.07-.77,6.91-.77,5.38,4.61-.77,4.61-2.3,6.14-9.98,8.45-29.95,7.68-49.15,15.36-82.18,32.26,35.33-9.98,58.37-11.52,73.73-12.29,7.68,0,10.75,0,9.98,4.61,0,3.84-.77,6.14-7.68,7.68-21.5,6.14-48.39,11.52-70.66,20.74,9.98-1.54,19.97-3.07,25.34-2.3,4.61.77,6.91,1.54,6.14,5.38-.77,4.61-3.84,6.91-9.22,7.68-19.2,1.54-25.34,3.84-48.39,10.75,4.61,0,4.61,3.07,2.3,5.38-2.3,3.07-5.38,5.38-9.22,5.38-6.91,0-21.5.77-32.26,7.68-6.14,4.61-9.22-6.14-6.14-9.98,8.45-9.22,33.79-23.04,46.85-32.26-25.34,12.29-49.15,23.04-71.43,43.78-7.68,6.91-11.52,3.84-12.29.77-.77-3.84-.77-6.14-4.61-3.84-3.07,1.54-4.61,3.07-9.98,7.68h0l-6.14,6.14h0l-32.26,29.18v-6.91l25.34-23.04v-6.91l-14.59,13.06v-6.91l21.5-19.2v-18.43l17.66-1.54-13.82-5.38c7.68-6.14,14.59-1.54,16.9.77,0-.77.77-6.14-3.84-9.22-2.3-.77-4.61-2.3-9.22-1.54-6.14-5.38-13.06-2.3-16.13.77-5.38-6.91-12.29-7.68-12.29-7.68l5.38,11.52-9.22-2.3,6.91,9.22c-3.84,6.14-8.45,13.82-12.29,27.65-1.54,5.38-3.07,16.13-3.07,26.11v2.3c-8.45-19.2-17.66-49.92-17.66-86.02,0-23.04,3.84-56.07,20.74-91.4,17.66-37.63,39.17-58.37,55.3-72.96,7.68-6.91,11.52-9.22,16.13-9.98,3.07,0,5.38.77,6.91,4.61,2.3,7.68,1.54,8.45,0,11.52-4.61,7.68-18.43,23.04-25.34,43.01,14.59-16.9,13.82-15.36,23.04-23.81,2.3-2.3,6.14-4.61,7.68,0,2.3,5.38,2.3,6.14-.77,10.75-2.3,3.84-13.82,14.59-26.11,37.63,6.14-6.91,9.22-10.75,10.75-11.52,2.3-2.3,4.61-5.38,7.68-.77,3.84,5.38,1.54,5.38,0,9.22-.77,3.07-13.06,22.27-21.5,38.4,6.14-6.91,6.91-8.45,9.98-9.98,1.54-.77,3.07,0,3.84,2.3s.77,6.14,0,7.68c-7.68,16.13-12.29,18.43-20.74,46.85,6.91-8.45,7.68-9.22,10.75-13.06,2.3-1.54,5.38-4.61,7.68,0,1.54,2.3,2.3,5.38.77,9.98-3.84,9.22-4.61,10.75-12.29,26.88,2.3-.77,3.07,0,4.61,2.3,6.14-7.68,13.06-14.59,18.43-19.97,20.74-21.5,39.94-38.4,69.12-57.6,33.03-22.27,60.67-35.33,79.87-43.01,9.98-3.07,17.66-6.14,21.5-6.14,3.84-.77,6.91,0,6.91,3.07s-.77,5.38-2.3,6.14c-5.38,3.07-9.22,4.61-23.04,11.52,21.5-6.14,47.62-9.22,59.14-7.68M1735.63,500.77c1.54-6.14,5.38-18.43,12.29-32.26-7.68,9.98-11.52,17.66-15.36,27.65,1.54.77,2.3,2.3,3.07,4.61M1711.06,493.09c8.45-29.18,23.81-82.18,49.15-119.81-17.66,19.2-43.78,62.21-49.15,119.81M1807.06,509.22c21.5-16.9,89.09-62.98,138.24-80.64-58.37,14.59-104.45,48.39-138.24,80.64M1874.65,587.55c-3.07-8.45-11.52-13.82-22.27-13.06-6.14-11.52-16.13-6.14-19.2-3.07,3.07,3.07,5.38,3.84,9.22,4.61v3.07c-6.91,0-9.22-3.07-13.82-8.45-2.3,16.9,17.66,23.04,23.81,9.22,9.22,0,15.36,3.84,16.9,9.98,2.3,6.14,0,12.29-3.07,15.36-5.38,6.91-14.59,6.91-23.04,2.3-5.38-2.3-10.75-7.68-15.36-13.06-6.91-7.68-10.75-11.52-16.13-13.82-3.07-1.54-6.14-2.3-8.45-2.3-1.54,0-3.07-.77-5.38-.77-1.54,0-3.07,0-4.61.77-12.29,1.54-24.58,6.91-39.17,8.45h-7.68c-4.61,4.61-9.22,8.45-13.06,11.52l-9.22,7.68h0c-3.84,3.84-6.14,5.38-6.14,5.38.77,9.22,1.54,13.82,1.54,21.5v12.29c0,3.07-.77,5.38-1.54,7.68s-2.3,3.84-5.38,3.84c-2.3.77-8.45-3.84-10.75-.77-1.54,2.3,0,7.68,2.3,9.22h16.9c2.3-4.61,3.07-6.91,3.84-9.22,1.54-2.3,2.3-5.38,3.84-6.91,1.54-2.3,2.3-4.61,3.07-6.91,0-2.3.77-3.84,1.54-5.38,1.54-3.07,3.07-6.14,3.84-7.68,0-.77.77-.77.77-1.54.77.77,1.54,2.3,2.3,3.07,3.84,6.14,7.68,10.75,13.06,13.06,2.3.77,3.84,2.3,4.61,3.07.77.77.77,3.07.77,5.38s-.77,4.61-3.07,5.38c-3.84,1.54-9.22-3.84-11.52-.77-2.3,2.3,0,6.14,2.3,9.22h15.36c3.84-6.14,5.38-13.82,5.38-20.74s0-8.45-4.61-10.75c-1.54-.77-3.07-4.61-3.84-7.68v-2.3c6.91-2.3,15.36-6.91,23.81-13.06v3.84c1.54,9.98,2.3,18.43,9.98,25.34,2.3,2.3,3.84,3.84,3.84,5.38,0,.77-.77,3.84-2.3,5.38-1.54,3.84-5.38,6.91-8.45,6.91-3.84.77-7.68-4.61-10.75-1.54-2.3,2.3.77,7.68,2.3,9.22h15.36c5.38-9.22,12.29-16.9,14.59-22.27,1.54-3.84,1.54-5.38-1.54-6.14-5.38-3.07-4.61-9.22-3.07-16.13,0-1.54.77-2.3,1.54-3.84,3.84,10.75,9.22,20.74,20.74,23.81,2.3,0,4.61,1.54,5.38,2.3v4.61c-.77,3.07-2.3,7.68-5.38,9.22-3.07,2.3-9.98-5.38-12.29-.77-1.54,3.07,1.54,7.68,2.3,9.22h15.36c3.84-9.22,7.68-17.66,9.22-23.81.77-6.91.77-9.98-4.61-10.75-3.07,0-5.38-1.54-6.14-5.38-2.3-6.91-1.54-19.2-6.91-26.88,6.14,6.91,12.29,12.29,17.66,15.36,13.06,6.91,26.88,4.61,33.03-4.61,3.84-9.22,3.84-16.9,2.3-22.27M1748.69,525.34c0,.77.77,2.3,2.3,2.3.77,0,2.3-.77,2.3-2.3,0-.77-.77-2.3-2.3-2.3s-2.3,1.54-2.3,2.3M1714.13,601.38v6.91l2.3-1.54c-.77-1.54-.77-3.84-1.54-5.38h-.77Z" />
                                </g>
                            </svg>
                        </tr>
                        <tr>
                            <td class="wrapper">
                                <p>
                                    The debugging proxy is <span style="font-weight: bold; color: green;">up and
                                        running</span>!
                                </p>

                                <p><strong>Connect your Rover</strong><br /> Use the <i>transceiver</i> service to
                                    connect
                                    to this debugging proxy server. Update the <i>service.yaml</i> and set the value of
                                    "passthrough-address" to
                                    the address of this passthrough server. 
                                </p>

                                <p><strong>View debug information</strong><br /> Use the <i>web-monitor</i> and enter
                                    the address of the proxy server to connect.</p>

                                <p>For more information, go to the <a href="https://docs.ase.vu.nl">ASE docs page</a>.
                                </p>
                            </td>
                        </tr>
                        <!-- END MAIN CONTENT AREA -->
                    </table>


                    <!-- END CENTERED WHITE CONTAINER -->
                </div>
            </td>
            <td>&nbsp;</td>
        </tr>
    </table>
</body>

</html>
		`))
	})

	// To retrieve an SDP offer (and send back an SDP answer)
	http.HandleFunc("/client/sdp", JSONEndpoint("[💻 CLIENT ONLY]: Send your SDP offer as a JSON object", func(w http.ResponseWriter, r *http.Request) ([]byte, error) {

		// Parse offer from request body
		request := rtc.RequestSDP{}
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			return nil, err
		}

		// Process offer
		return events.OnClientSDPReceived(request, s)
	}))

	// To retrieve an ICE candidate (and send back an ICE candidate)
	http.HandleFunc("/client/ice", JSONEndpoint("[💻 CLIENT ONLY]: Send your ICE candidate as a JSON object", func(w http.ResponseWriter, r *http.Request) ([]byte, error) {
		// Parse ICE from request body
		request := rtc.RequestICE{}
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			return nil, err
		}

		return events.OnClientICEReceived(request, s)
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

		return events.OnCarSDPReceived(request, receivedAt, s)
	}))

	// To retrieve an ICE candidate (and send back an ICE candidate)
	http.HandleFunc("/car/ice", JSONEndpoint("[🚗 CAR ONLY]: Send your ICE candidate as a JSON object", func(w http.ResponseWriter, r *http.Request) ([]byte, error) {
		// Parse ICE from request body
		request := rtc.RequestICE{}
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			return nil, err
		}

		return events.OnCarICEReceived(request, s)
	}))

	// Start HTTP server to accept incoming connections
	log.Debug().Msgf("passthrough HTTP listener active on '%s'", serverAddress)

	err := http.ListenAndServe(serverAddress, nil)
	if err != nil {
		return fmt.Errorf("Cannot start HTTP server: %v", err)
	}
	return err
}
