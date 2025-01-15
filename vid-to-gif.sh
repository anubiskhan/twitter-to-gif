#!/bin/bash

# Check if URL is provided
if [ -z "$1" ]; then
    echo "Please provide a Twitter/X.com URL"
    echo "Usage: ./vid-to-gif.sh <url>"
    exit 1
fi

input_url="$1"

# Add random delay between 1-3 seconds
sleep $(( ( RANDOM % 3 ) + 1 ))

# Build the image if it doesn't exist
if ! docker image inspect vid-to-gif >/dev/null 2>&1; then
    echo "Building Docker image..."
    docker build -t vid-to-gif .
fi

# Run the container with proper error handling and rate limiting
docker run --rm -v ~/Downloads:/workdir vid-to-gif \
    sh -c "set -e && \
    echo 'Attempting to fetch video URL...' && \
    video_url=\$(yt-dlp --verbose -g '$input_url' -f 'bestvideo') && \
    if [ -z \"\$video_url\" ]; then exit 1; fi && \
    sleep 2 && \
    filename=\$(yt-dlp --get-filename -o '%(uploader)s_%(id)s' '$input_url') && \
    if [ -z \"\$filename\" ]; then filename=\"output\"; fi && \
    sleep 1 && \
    ffmpeg -i \"\$video_url\" \
    -vf 'fps=10,scale=480:-1:flags=lanczos' \
    -c:v gif \
    \"/workdir/\${filename}.gif\" && \
    echo \"Converted successfully to \${filename}.gif\"" || \
    { echo "Error: Failed to convert video"; exit 1; }

if [ $? -eq 0 ]; then
    echo "GIF saved to ~/Downloads/${filename}.gif"
else
    echo "Conversion failed"
    exit 1
fi