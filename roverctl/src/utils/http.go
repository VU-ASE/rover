package utils

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"
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
			return fmt.Errorf("%s", PrettyJSON(httpRes))
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
