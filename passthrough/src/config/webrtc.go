package livestreamconfig

const (
	// Used to identify the car connection
	CarId = "car"

	// Used to identify the different data channels
	ControlChannelLabel = "control"
	DataChannelLabel    = "data"

	// The UDP port to use for ICE candidate multiplexing
	// Updating this value also requires updating your Dockerfile and docker-compose.yaml
	MuxUdpPort = 40000
)
