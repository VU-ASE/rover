package configuration

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"slices"
	"time"

	"github.com/VU-ASE/roverctl/src/openapi"
	"github.com/melbahja/goph"
	"golang.org/x/crypto/ssh"
	"gopkg.in/yaml.v3"
)

//
// Get, create and set all the available connections to roverd endpoints
//

// The file name in the configuration directory where the connections are stored
var connectionsFileName = LocalConfigDir() + "/connections.yaml"

type RoverConnection struct {
	Name     string             `yaml:"name"`
	Host     string             `yaml:"host"`
	Username string             `yaml:"username"`
	Password string             `yaml:"password"`
	client   *openapi.APIClient // to be used to communicate with the roverd endpoint
}

// An overview of all the available connections, as is written to the configuration file
type RoverConnections struct {
	Available []RoverConnection `yaml:"available"`
	Active    string            `yaml:"active"`
}

// To read state from disk
func ReadConnections() (RoverConnections, error) {
	connections := RoverConnections{
		Available: []RoverConnection{},
		Active:    "",
	}

	// Check if the file exists
	if _, err := os.Stat(connectionsFileName); os.IsNotExist(err) {
		// If the file does not exist, return an empty array
		return connections, nil
	}

	// Read the file
	content, err := os.ReadFile(connectionsFileName)
	if err != nil {
		return connections, err
	}

	// Parse the YAML content
	err = yaml.Unmarshal(content, &connections)
	return connections, err
}

func (c RoverConnections) GetActive() *RoverConnection {
	for _, connection := range c.Available {
		if connection.Name == c.Active {
			return &connection
		}
	}
	return nil
}

func (c RoverConnections) Save() error {
	// Marshal the connections to YAML
	content, err := yaml.Marshal(c)
	if err != nil {
		return err
	}

	// Write the file, overwriting the existing one
	return os.WriteFile(connectionsFileName, content, 0644)
}

func (c RoverConnections) Add(new RoverConnection) RoverConnections {
	// If a connection with the same name already exists, remove it
	c.Available = slices.DeleteFunc(c.Available, func(c RoverConnection) bool {
		return c.Name == new.Name
	})

	c.Available = append(c.Available, new)
	c.Active = new.Name
	return c
}

func (c RoverConnections) Get(name string) *RoverConnection {
	for _, connection := range c.Available {
		if connection.Name == name {
			return &connection
		}
	}
	return nil
}

func (c RoverConnections) Remove(name string) RoverConnections {
	// Find the connection to remove
	c.Available = slices.DeleteFunc(c.Available, func(c RoverConnection) bool {
		return c.Name == name
	})
	// Set the active connection to the first one in the list
	if len(c.Available) > 0 {
		c.Active = c.Available[0].Name
	}
	return c
}

func (c RoverConnections) SetActive(name string) RoverConnections {
	// Check if the connection exists
	found := false
	for _, c := range c.Available {
		if c.Name == name {
			found = true
			break
		}
	}

	if found {
		c.Active = name
	}

	return c
}

// Convert the RoverConnection to an SSH connection object
// Don't forget to close!
func (c RoverConnection) ToSshConnection() (*ssh.Client, error) {
	config := &ssh.ClientConfig{
		User: c.Username,
		Auth: []ssh.AuthMethod{
			ssh.Password(c.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	return ssh.Dial("tcp", c.Host, config)
}

// Convert the RoverConnection to a goph SSH connection object (which often is more useful)
// Don't forget to close!
func (c RoverConnection) ToSsh() (*goph.Client, error) {
	auth := goph.Password(c.Password)
	return goph.NewConn(&goph.Config{
		User:     c.Username,
		Addr:     c.Host,
		Auth:     auth,
		Timeout:  goph.DefaultTimeout,
		Callback: ssh.InsecureIgnoreHostKey(),
	})
}

// To add Basic Auth headers
type authTransport struct {
	Username string
	Password string
	Base     http.RoundTripper
}

func (a *authTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Add Basic Authentication header
	auth := a.Username + ":" + a.Password
	encodedAuth := base64.StdEncoding.EncodeToString([]byte(auth))
	req.Header.Add("Authorization", "Basic "+encodedAuth)
	return a.Base.RoundTrip(req)
}

func basicAuthClient(username, password string) *http.Client {
	return &http.Client{
		Transport: &authTransport{
			Username: username,
			Password: password,
			Base:     http.DefaultTransport, // Use the default RoundTripper
		},
		Timeout: 600 * time.Second,
	}
}

func (c RoverConnection) ToApiClient() *openapi.APIClient {
	if c.client == nil {
		config := openapi.NewConfiguration()
		config.Servers = openapi.ServerConfigurations{
			{
				URL: fmt.Sprintf("http://%s", c.Host),
			},
		}
		config.HTTPClient = basicAuthClient(c.Username, c.Password)
		c.client = openapi.NewAPIClient(config)
	}
	return c.client
}
