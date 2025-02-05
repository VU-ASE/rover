#!/bin/bash

set -e

REPO="VU-ASE/roverctl"
INSTALL_DIR="/usr/local/bin"

# Function to detect OS
detect_os() {
  uname_out="$(uname -s)"
  case "${uname_out}" in
    Linux*) os="linux";;
    Darwin*) os="macos";;
    *) echo "Unsupported OS: ${uname_out}"; exit 1;;
  esac
}

# Function to detect architecture
detect_arch() {
  uname_arch="$(uname -m)"
  case "${uname_arch}" in
    x86_64) arch="amd64";;
    arm64|aarch64) arch="arm64";;
    *) echo "Unsupported architecture: ${uname_arch}"; exit 1;;
  esac
}

# Function to determine the shell's profile for PATH setup
detect_shell_profile() {
  case "$SHELL" in
    */bash) echo "$HOME/.bashrc";;
    */zsh) echo "$HOME/.zshrc";;
    */fish) echo "$HOME/.config/fish/config.fish";;
    *) echo "Unknown shell. Please manually add $INSTALL_DIR to your PATH."; exit 1;;
  esac
}

# Detect OS and architecture
detect_os
detect_arch

# Construct the binary name
binary_name="roverctl-${os}-${arch}"

# Get the latest release tag
echo "Fetching the latest release..."
latest_release=$(curl -s "https://api.github.com/repos/$REPO/releases/latest" | sed -n 's/.*"tag_name": *"\([^"]*\)".*/\1/p')

if [ -z "$latest_release" ]; then
  echo "Failed to fetch the latest release from $REPO."
  exit 1
fi

echo "Latest release: $latest_release"

# Download the binary
echo "Downloading the binary for ${os}/${arch}..."
curl -Lo "/tmp/${binary_name}" "https://github.com/${REPO}/releases/download/${latest_release}/${binary_name}"

# Make it executable
chmod +x "/tmp/${binary_name}"

# Move the binary to the install directory
echo "Installing the binary to ${INSTALL_DIR}..."
sudo mv "/tmp/${binary_name}" "${INSTALL_DIR}/roverctl"

# Add to PATH if necessary
if ! command -v roverctl &> /dev/null; then
  echo "roverctl is not in your PATH. Attempting to add it..."
  shell_profile=$(detect_shell_profile)
  echo "export PATH=\"${INSTALL_DIR}:\$PATH\"" >> "$shell_profile"
  echo "Added ${INSTALL_DIR} to PATH in $shell_profile. Please restart your shell or run 'source $shell_profile'."
fi

echo "Installation complete! You can now use 'roverctl'."
