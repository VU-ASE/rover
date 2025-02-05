#!/bin/bash

if [ $# -ne 0 ]; then
    echo "Usage: cat CHANGELOG.md | $0"
    exit 1
fi


# Read the input and get the first line that contains a version number
# Using grep with brackets to find version patterns, then taking first line with head -n1
first_version=$(grep -o '\[[0-9]\+\.[0-9]\+\.[0-9]\+\]' | head -n1)

# Remove the square brackets from the version number using sed
# sed 's/pattern/replacement/' removes brackets by replacing [version] with just version
clean_version=$(echo "$first_version" | sed 's/\[\(.*\)\]/\1/')

# Output the clean version number
echo "$clean_version"
