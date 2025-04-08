# Installation

The `roverctl` utility runs on your own system. You can conveniently install it using our pre-built binaries or build from source using our `Makefile` and provided devcontainer.

## Install pre-built binaries (recommended)

Linux and macOS users (both amd64 and arm64) can install our pre-built binaries using the provided installation script. This script will detect your system and add `roverctl` to your `PATH` automatically:

```bash
# Install latest
curl -fsSL https://raw.githubusercontent.com/VU-ASE/rover/refs/heads/main/roverctl/install.sh | bash
# Install a specific version (i.e. 1.0.0)
curl -fsSL https://raw.githubusercontent.com/VU-ASE/rover/refs/heads/main/roverctl/install.sh | bash -s v1.0.0
```

Alternatively, you can download the pre-built binaries and releases [here](https://github.com/VU-ASE/rover/releases/latest).

## Self-updating

If you already have `roverctl` installed you can update to the latest version directly:

```bash
# Install latest
roverctl update roverctl
# Install specific version (i.e. 1.0.0)
roverctl update roverctl --version 1.0.0
```

## Build from source

If you cannot or do not want to install the pre-built binaries (see above), you can build `roverctl` from source.

To install the repository from source, you can use our Makefile:
```bash
git clone https://github.com/VU-ASE/rover.git
cd rover/roverctl
make build
# Run roverctl from the build directory (not in PATH yet)
./bin/roverctl
```

We provide users with a *.devcontainer* that can be used in VS Code and has all necessary dependencies installed already. If you want to understand which dependencies need to be installed, take a look at the [*.devcontainer/Dockerfile*](https://github.com/VU-ASE/rover/blob/main/.devcontainer/roverctl/Dockerfile).
