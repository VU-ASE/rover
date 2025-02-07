#!/bin/bash

# This script should be used only as part of the development process for a developer
# to have a setup that is identical to the rovers. The key files/directories are:
# /etc/roverd/rover.yaml
# /etc/roverd/info.txt
# /home/debix/.rover

handle_error() {
    local cmd=$1
    local exit_code=$2
    echo "Error: Failed to execute: $cmd (exit code: $exit_code)" >&2
    exit -1
}

PROJECT_ROOT=/workspace/rover/roverd/
TEST_FILES=$PROJECT_ROOT/rovervalidate/src/testfiles
ROVER_INFO_FILE=/etc/roverd/info.txt

# --- /etc/roverd/rover.yaml ---
mkdir -p /etc/roverd || handle_error "mkdir -p /etc/roverd" $?
sudo cp $TEST_FILES/roverd-yaml/valid/empty.yaml /etc/roverd/rover.yaml || handle_error "cp empty.yaml" $?

# Create debix user if it doesn't exist
id -u debix &>/dev/null || useradd -m -s /bin/bash debix || handle_error "useradd debix" $?

# --- /home/debix/rover/ ---
mkdir -p /home/debix/.rover || handle_error "mkdir -p /home/debix/.rover" $?
chown debix:debix /home/debix/.rover || handle_error "chown debix:debix" $?

# --- /etc/rover ---
echo "14" > $ROVER_INFO_FILE || handle_error "write rover id" $?
echo "bunny" >> $ROVER_INFO_FILE || handle_error "write rover name" $?
echo -n "debix" | sha256sum | cut -d ' ' -f 1 >> $ROVER_INFO_FILE || handle_error "write password hash" $?