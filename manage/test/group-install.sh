#!/bin/bash

# Define the IP range and corresponding rover names
start_ip=1
end_ip=20
base_ip="192.168.0."

# Create a temporary directory for storing results
tmp_dir=$(mktemp -d)
trap 'rm -rf "$tmp_dir"' EXIT

# Function to check a single rover and store results
check_rover() {
    local i=$1
    local ip="$base_ip$i"
    local rover_index=$((i - start_ip + 1))
    local rover_host_name=$(printf "rover%02d.local" "$rover_index")
    local result_file="$tmp_dir/$i.result"

    ./all.sh "$i" safe install
}

# Launch all checks in parallel
for ((i=$start_ip; i<=$end_ip; i++)); do
    check_rover "$i" &
done

echo "Fetching status of all rovers"

# Wait for all background processes to complete
wait
