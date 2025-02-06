# Installation

The `roverctl` utility runs on your own system. You can conveniently install it using our pre-built binaries or build from source using our `Makefile` and provided devcontainer.

## Install pre-built binaries (recommended)

Linux and macOS users (both amd64 and arm64) can install our pre-built binaries using our installation script. This script will detect your system and add `roverctl` to your `PATH` automatically:

```bash
# Install latest
curl -fsSL https://raw.githubusercontent.com/VU-ASE/rover/refs/heads/main/roverctl/install.sh | bash
# Install a specific version (i.e. 1.0.0)
curl -fsSL https://raw.githubusercontent.com/VU-ASE/rover/refs/heads/main/roverctl/install.sh | bash -s v1.0.0
```

Alternatively, you can download the pre-built binaries and releases [here](https://github.com/VU-ASE/rover/releases/latest).

## Build from source

To install the repository from source, you can use our Makefile:
```bash
git clone https://github.com/VU-ASE/rover.git
cd rover/roverctl
make build
# Run roverctl from the build directory (not in PATH yet)
./bin/roverctl
```

We provide users with a *.devcontainer* that can be used in VS Code and has all necessary dependencies installed already. If you want to understand which dependencies need to be installed, take a look at the *.devcontainer/Dockerfile*.