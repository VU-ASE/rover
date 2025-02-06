#!/bin/bash

# Define the IP range and corresponding rover names
start_ip=101
end_ip=120
base_ip="192.168.0."

# Function to display a green checkmark
function green_check {
    printf "\033[32m✔\033[0m"
}

# Function to display an orange warning
function orange_warning {
    printf "\033[33m⚠\033[0m"
}

# Function to display a red cross
function red_cross {
    printf "\033[31m✘\033[0m"
}

# Read the credentials from the rovers.txt file
credentials=( $(cat rovers.txt) )

# Iterate over the range of IPs
for ((i=$start_ip; i<=$end_ip; i++)); do
    ip="$base_ip$i"
    rover_index=$((i - start_ip + 1))
    rover_name=$(printf "rover%02d" "$rover_index")

    # Extract username and password for the current rover
    creds_index=$((rover_index - 1))
    username=$(echo "${credentials[$creds_index]}" | cut -d';' -f1)
    password=$(echo "${credentials[$creds_index]}" | cut -d';' -f2)

    # Check if the IP is reachable via ping
    if ping -c 1 -W 2 "$ip" > /dev/null 2>&1; then
        # Make the POST request with BasicAuth
        response=$(curl -s -m 2 -w "\n%{http_code}" -u "$username:$password" -X POST "http://$ip/pipeline/stop")
        
        # Extract HTTP status code
        http_status=$(echo "$response" | tail -n1)

        # Check the HTTP status code
        if [[ $http_status =~ ^2 ]]; then
            printf "%s: " "$rover_name"
            green_check
            printf " - POST request successful\n"
        elif [[ $http_status =~ ^[3-5] ]]; then
            printf "%s: " "$rover_name"
            orange_warning
            printf " - POST request returned status code %s\n" "$http_status"
        else
            printf "%s: " "$rover_name"
            red_cross
            printf " - POST request failed\n"
        fi
    else
        printf "%s: " "$rover_name"
        red_cross
        printf "\n"
    fi

done
