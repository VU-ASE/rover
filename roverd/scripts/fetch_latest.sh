#!/bin/bash


ROVER_IP=localhost
ROVER_PW=debix

function check_roverd {
    status=`curl -s $ROVER_IP/status | jq -r '.status'`

    if [ "$status" != "operational" ] ; then
        echo Error: roverd not operational;
        exit 1;
    fi
}

function get_service {
    curl -u debix:$ROVER_PW \
        -X 'POST' \
        "http://$ROVER_IP/fetch" \
        -H 'accept: application/json' \
        -H 'Content-Type: application/json' \
        -d "{\"url\": \"https://github.com/VU-ASE/$1/releases/latest/download/$1.zip\"}"
}

check_roverd

get_service imaging
get_service controller
get_service actuator



curl -u debix:debix \
        -X 'POST' \
        "http://localhost/fetch" \
        -H 'accept: application/json' \
        -H 'Content-Type: application/json' \
        -d "{\"url\": \"https://github.com/VU-ASE/$1/releases/latest/download/$1.zip\"}"

curl -u debix:debix \
    -X 'POST' \
    "http://192.168.0.112/services/vu-ase/imaging/1.2.1/configuration" \
    -H 'accept: application/json' \
    -H 'Content-Type: application/json' \
    -d "[{\"key\": \"gstreamer-pipeline\", \"value\": \"haram\"}]"



