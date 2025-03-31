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
    echo "Building normal pipeline"
    echo "Nothing to build"

elif [[ "$mode" == "safe" ]]; then
    echo "Building safe pipeline"

    # Build safe controller
    roverctl service build tester safe-controller 1.0.0 -r "$number"
fi