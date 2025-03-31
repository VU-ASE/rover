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

roverctl update roverd -r "$number"

if [[ "$mode" == "normal" ]]; then
    echo "Installing normal pipeline"
    roverctl author --set tester

    # Install imaging 
    roverctl service install https://github.com/VU-ASE/imaging/releases/download/v1.2.4/imaging.zip -r "$number"
    # Install controller
    roverctl service install https://github.com/VU-ASE/controller/releases/download/v1.4.2/controller.zip -r "$number"
    # Install actuator
    roverctl service install https://github.com/VU-ASE/actuator/releases/download/v1.3.2/actuator.zip -r "$number"

elif [[ "$mode" == "safe" ]]; then
    echo "Installing safe pipeline"

    roverctl author --set tester
    
    # Install safe controller
    roverctl upload ../safe-controller -r "$number"
    # Install distance sensor
    roverctl service install https://github.com/VU-ASE/distance/releases/download/v1.0.2/distance.zip -r "$number"
    # Install imaging 
    roverctl service install https://github.com/VU-ASE/imaging/releases/download/v1.2.4/imaging.zip -r "$number"
    # Install actuator
    roverctl service install https://github.com/VU-ASE/actuator/releases/download/v1.3.2/actuator.zip -r "$number"
fi