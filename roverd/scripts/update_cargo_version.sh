#!/bin/bash

# Check if a version argument was provided
if [ $# -ne 2 ]; then
    echo "Usage: $0 <version-number> <Cargo.toml-path>"
    echo "Example: $0 1.1.5 ./roverd/Cargo.toml"
    exit 1
fi

# Store the version number from command line argument
VERSION=$1

# Cargo toml file path
TOML_FILE=$2

# Validate that the version matches semantic versioning format (x.y.z)
if ! [[ $VERSION =~ ^[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
    echo "Error: Version must be in format x.y.z (e.g., 1.1.5)"
    exit 1
fi

# Check if Cargo.toml exists in the current directory
if [ ! -f "$TOML_FILE" ]; then
    echo "Error: $TOML_FILE not found in current directory"
    exit 1
fi

# Create a backup of the original file
cp $TOML_FILE $TOML_FILE.backup

# Use sed to replace the version line
# The pattern looks for 'version = "x.y.z"' and replaces it with the new version
# The -i flag means "in-place" editing
# The -E flag enables extended regular expressions
sed -i.tmp -E "s/^version = \"[0-9]+\.[0-9]+\.[0-9]+\"/version = \"$VERSION\"/" $TOML_FILE

# Check if the replacement was successful
if grep -q "version = \"$VERSION\"" $TOML_FILE; then
    echo "Successfully updated version to $VERSION"
    # Remove the temporary file created by sed
    rm $TOML_FILE.tmp
else
    echo "Error: Failed to update version. Restoring backup..."
    mv $TOML_FILE.backup $TOML_FILE
    exit 1
fi

# Remove the backup file if everything went well
rm $TOML_FILE.backup