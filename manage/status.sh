#!/bin/bash

# Define the IP range and corresponding rover names
start_ip=101
end_ip=120
base_ip="192.168.0."

# Function to display a green checkmark
function green_check {
    printf "\033[32m✔\033[0m"
}

# Function to display a red cross
function red_cross {
    printf "\033[31m✘\033[0m"
}

# Function to display an orange warning
function orange_warning {
    printf "\033[33m⚠\033[0m"
}

# Create a temporary directory for storing results
tmp_dir=$(mktemp -d)
trap 'rm -rf "$tmp_dir"' EXIT

# Function to check a single rover and store results
check_rover() {
    local i=$1
    local ip="$base_ip$i"
    local rover_index=$((i - start_ip + 1))
    local rover_name=$(printf "rover%02d" "$rover_index")
    local result_file="$tmp_dir/$i.result"

    # Check if the IP is reachable via ping
    if ping -c 1 -W 2 "$ip" > /dev/null 2>&1; then
        response=$(curl -s -m 0.1 -w "\n%{http_code}" -X GET "http://$ip/status")
        
        # Extract HTTP status code
        http_status=$(echo "$response" | tail -n1)

        # Check if the HTTP status code is 2XX
        if [[ $http_status =~ ^2 ]]; then
            # Extract "rover_name" from the JSON response
            rover_response_name=$(echo "$response" | head -n1 | jq -r '.rover_name')
            roverd_version=$(echo "$response" | head -n1 | jq -r '.version')

            echo "success|$rover_name|$rover_response_name|$roverd_version" > "$result_file"
        else
            echo "warning|$rover_name" > "$result_file"
        fi
    else
        echo "error|$rover_name" > "$result_file"
    fi
}

# Launch all checks in parallel
for ((i=$start_ip; i<=$end_ip; i++)); do
    check_rover "$i" &
done

echo "Fetching status of all rovers"

# Wait for all background processes to complete
wait

# Sort and display results
for ((i=$start_ip; i<=$end_ip; i++)); do
    result_file="$tmp_dir/$i.result"
    if [[ -f "$result_file" ]]; then
        IFS='|' read -r status rover_name rover_response_name roverd_version < "$result_file"
        
        printf "%s: " "$rover_name"
        case "$status" in
            "success")
                green_check
                printf " roverd@$roverd_version - %s\n" "$rover_response_name"

                ;;
            "warning")
                orange_warning
                printf " - GET /status request failed\n"
                ;;
            "error")
                red_cross
                printf "\n"
                ;;
        esac
    fi
done