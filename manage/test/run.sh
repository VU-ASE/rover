#!/bin/bash

# Exit on error
set -e

# Ensure number is passed
if [[ -z "$1" || ! "$1" =~ ^[0-9]+$ || "$1" -lt 0 || "$1" -gt 20 ]]; then
    echo "Usage: $0 <number between 0 and 20> <mode: normal|safe>"
    exit 1
fi

# Ensure mode is passed and valid
if [[ -z "$2" || ( "$2" != "normal" && "$2" != "safe" ) ]]; then
    echo "Usage: $0 <number between 0 and 20> <mode: normal|safe>"
    exit 1
fi

number="$1"
mode="$2"

if [[ "$mode" == "normal" ]]; then
    echo "Running normal pipeline"

    roverctl pipeline reset -r "$number"

    # Enable all services (order is important)
    roverctl pipeline enable vu-ase imaging 1.2.4 -r "$number"
    roverctl pipeline enable vu-ase controller 1.4.2 -r "$number"
    roverctl pipeline enable vu-ase actuator 1.3.1 -r "$number"

    # Start the pipeline
    roverctl pipeline start -r "$number"
    # Query pipeline
    roverctl pipeline -r "$number"

elif [[ "$mode" == "safe" ]]; then
    echo "Running safe pipeline"
    roverctl pipeline reset -r "$number"

    # Enable all services (order is important)
    roverctl pipeline enable vu-ase distance 1.0.2 -r "$number"
    roverctl pipeline enable vu-ase imaging 1.2.4 -r "$number"
    roverctl pipeline enable tester safe-controller 1.0.0 -r "$number"
    roverctl pipeline enable vu-ase actuator 1.3.1 -r "$number"

    # Start the pipeline
    roverctl pipeline start -r "$number"

    # Query pipeline
    roverctl pipeline -r "$number"
fi

# Wait for any key to be pressed
echo "Press any key to stop the pipeline..."
read -n 1 -s
echo "Stopping pipeline..."
roverctl pipeline stop -r "$number"
