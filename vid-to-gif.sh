#!/bin/bash

# Check if URL is provided
if [ -z "$1" ]; then
    echo "Please provide a URL"
    echo "Usage: ./vid-to-gif.sh <url> [gif|video]"
    exit 1
fi

input_url="$1"
mode="${2:-gif}"  # Default to gif if not specified

# Validate mode
if [[ "$mode" != "gif" && "$mode" != "video" ]]; then
    echo "Invalid mode: $mode"
    echo "Mode must be either 'gif' or 'video'"
    exit 1
fi

# Get user's home directory and Downloads path
if [[ "$OSTYPE" == "darwin"* ]]; then
    # macOS
    DOWNLOADS_DIR="$HOME/Downloads"
else
    # Linux and others
    DOWNLOADS_DIR="${XDG_DOWNLOAD_DIR:-$HOME/Downloads}"
fi

# Create Downloads directory if it doesn't exist
mkdir -p "$DOWNLOADS_DIR"

# Build the images if they don't exist
echo "Building Docker images..."
docker build -t twitter-to-gif-app -f Dockerfile.app .
docker build -t twitter-to-gif-converter -f Dockerfile.converter .

# Run the app container
docker run --rm \
    -v "$DOWNLOADS_DIR:/workdir" \
    -v /var/run/docker.sock:/var/run/docker.sock \
    twitter-to-gif-app "$input_url" "$mode"