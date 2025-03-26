package utils

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/VU-ASE/rover/roverctl/src/openapi"
)

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
			content := PrettyJSON(httpRes)

			// If we can parse this as a generic error message, do so
			// and show a cleaner error message
			var roverdError openapi.RoverdError
			err = roverdError.UnmarshalJSON(httpRes)
			if err == nil {
				genericError := roverdError.ErrorValue.GenericError
				if genericError != nil {
					content = fmt.Sprintf("generic error, code %d\n", genericError.Code)
					content += genericError.Message
				}
			}

			// Try to actually print the \n characters
			unescaped := strings.ReplaceAll(content, `\n`, "\n")
			unescaped = strings.ReplaceAll(unescaped, `\"`, "\"")

			return fmt.Errorf("%s", unescaped)
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
