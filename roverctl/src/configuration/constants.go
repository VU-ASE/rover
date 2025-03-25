package configuration

const (
	ROVERCTL_UPDATE_LATEST_SCRIPT = "curl -fsSL https://raw.githubusercontent.com/VU-ASE/rover/refs/heads/main/roverctl/install.sh | bash"
	VU_ASE_AUTHOR                 = "vu-ase"
)

var (
	ROVERCTL_UPDATE_LATEST_SCRIPT_WITH_VERSION = "curl -fsSL https://raw.githubusercontent.com/VU-ASE/rover/refs/heads/main/roverctl/install.sh | bash -s " // ... followed by vX.Y.Z
)
