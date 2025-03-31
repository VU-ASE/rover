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

ip="192.168.0.1$(printf "%02d" "$number")"  # Converts 6 -> 06, gives 192.168.0.106

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

# Temp dir for results
tmp_dir=$(mktemp -d)
trap 'rm -rf "$tmp_dir"' EXIT

# Check rover function
check_rover() {
    local ip="$1"
    local result_file="$tmp_dir/result"

    if ping -c 1 -W 2 "$ip" > /dev/null 2>&1; then
        response=$(curl -s -m 0.1 -w "\n%{http_code}" -X GET "http://$ip/status")
        http_status=$(echo "$response" | tail -n1)

        if [[ $http_status =~ ^2 ]]; then
            rover_name=$(echo "$response" | head -n1 | jq -r '.rover_name')
            roverd_version=$(echo "$response" | head -n1 | jq -r '.version')
            echo "success|$rover_name|$roverd_version" > "$result_file"
        else
            echo "warning|$ip" > "$result_file"
        fi
    else
        echo "error|$ip" > "$result_file"
    fi
}

# Run the check
check_rover "$ip"

# Parse result
IFS="|" read -r status rover_info extra <<< "$(cat "$tmp_dir/result")"

case "$status" in
    success)
        echo -n "[$ip] "
        green_check
        echo " - Rover: $rover_info, Version: $extra"

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



 
        ;;
    warning)
        echo -n "[$ip] "
        orange_warning
        echo " - Warning: Non-2xx HTTP response"
        ;;
    error)
        echo -n "[$ip] "
        red_cross
        echo " - Error: Host unreachable"
        ;;
esac
