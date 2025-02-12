package utils

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/VU-ASE/rover/roverctl/src/openapi"
)

func ParseRoverdError(json []byte) *openapi.RoverdError {
	// Try to parse as error first
	roverdErr := &openapi.RoverdError{}
	err := roverdErr.UnmarshalJSON([]byte(json))
	if err != nil {
		return nil
	}

	return roverdErr
}

// Convert a roverd error to a colored string that can be used in the TUI
// If not a valid error, its string representation is returned
func RoverdErrorToString(rd openapi.RoverdError) string {
	switch {
	case rd.ErrorValue.PipelineSetError != nil:
		rdps := rd.ErrorValue.PipelineSetError
		s := "Setting pipeline failed"
		for _, v := range rdps.ValidationErrors.UnmetStreams {
			s += fmt.Sprintf("\nService %s depends on stream %s from service %s, but this stream was not published.", v.Source, v.Stream, v.Target)
		}
		return s

	case rd.ErrorValue.GenericError != nil:
		return fmt.Sprintf("Generic error: (%d) %v\n", rd.ErrorValue.GenericError.Code, rd.ErrorValue.GenericError.Message)
	default:
		return fmt.Sprintf("Unknown error: %v\n", rd)
	}
}

func ParseHTTPError(err error, htt *http.Response) error {
	if err != nil && htt != nil {
		var netErr net.Error
		if errors.As(err, &netErr) && netErr.Timeout() {
			return fmt.Errorf("Operation timed out")
		}

		// Read http response body
		httpRes := make([]byte, htt.ContentLength)
		_, err = htt.Body.Read(httpRes)
		if err != nil {
			return fmt.Errorf("Failed to read http response body: %v", err)
		} else {
			rd := ParseRoverdError(httpRes)
			if rd != nil {
				return fmt.Errorf("%s", RoverdErrorToString(*rd))
			} else {
				return fmt.Errorf("Unknown error: %s\n", string(httpRes))
			}
		}
	} else {
		return fmt.Errorf("Operation failed: %v", err)
	}
}

func FollowRedirects(url string) (string, error) {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			// Allow following redirects up to a reasonable limit
			if len(via) > 10 {
				return fmt.Errorf("too many redirects")
			}
			return nil
		},
	}

	// Make the request
	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Extract the final URL
	finalURL := resp.Request.URL.String()
	return finalURL, nil
}

func IsHostOnline(host string, port string, timeout time.Duration) bool {
	address := net.JoinHostPort(host, port)
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}
