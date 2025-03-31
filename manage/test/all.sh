#!/bin/bash

# Exit on error
set -e

# Ensure a number is passed
if [[ -z "$1" || ! "$1" =~ ^[0-9]+$ || "$1" -lt 0 || "$1" -gt 20 ]]; then
    echo "Usage: $0 <number between 0 and 20>"
    exit 1
fi

# Ensure mode is passed and valid
if [[ -z "$2" || ( "$2" != "normal" && "$2" != "safe" ) ]]; then
    echo "Usage: $0 <number between 0 and 20> <mode: normal|safe>"
    exit 1
fi

# Ensure method is passed and valid
if [[ -z "$3" || ( "$3" != "install" && "$3" != "build" && "$3" != "run" ) ]]; then
    echo "Usage: $0 <number between 0 and 20> <mode: install|build|run> <method: install|build|run>"
    exit 1
fi

# Variables
number="$1"
mode="$2"
method="$3"

ip="rover$(printf "%02d" "$number").local"  # Converts 6 -> 06

echo "Testing $ip"

# Icon functions
function green_check {
    printf "\033[32m✔\033[0m"
}

function red_cross {
    printf "\033[31m✘\033[0m"
}

function orange_warning {
    printf "\033[33m⚠\033[0m"
}


  # If method is install, install first, then run the pipeline
        if [[ "$method" == "install" ]]; then
            echo "Installing the pipeline..."
            ./install.sh "$number" "$mode" 
        elif [[ "$method" == "build" ]]; then
            echo "Building the pipeline..."
            ./build.sh "$number" "$mode"
        elif [[ "$method" == "run" ]]; then
            echo "Running the pipeline..."
            ./run.sh "$number" "$mode"
        fi