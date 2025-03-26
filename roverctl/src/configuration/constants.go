package configuration

const (
	ROVERCTL_UPDATE_LATEST_SCRIPT = "curl -fsSL https://raw.githubusercontent.com/VU-ASE/rover/refs/heads/main/roverctl/install.sh | bash"
	VU_ASE_AUTHOR                 = "vu-ase"

	// Used to identify the car connection
	PROXY_CAR_ID = "car"

	// Used to identify the different data channels
	PROXY_CONTROL_CHAN_LABEL = "control"
	PROXY_DATA_CHAN_LABEL    = "data"
)

var (
	ROVERCTL_UPDATE_LATEST_SCRIPT_WITH_VERSION = "curl -fsSL https://raw.githubusercontent.com/VU-ASE/rover/refs/heads/main/roverctl/install.sh | bash -s " // ... followed by vX.Y.Z
)
