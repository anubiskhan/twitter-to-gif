#!/bin/bash

# Check if URL is provided
if [ -z "$1" ]; then
    echo "Please provide a Twitter/X.com URL"
    echo "Usage: ./vid-to-gif.sh <url>"
    exit 1
fi

input_url="$1"

# Add random delay between 0.1-0.3 seconds
sleep $(awk -v min=0.1 -v max=0.3 'BEGIN{srand(); print min+rand()*(max-min)}')

# Build the image if it doesn't exist
if ! docker image inspect vid-to-gif >/dev/null 2>&1; then
    echo "Building Docker image..."
    docker build -t vid-to-gif .
fi

# Run the container with GIF detection and handling
docker run --rm -v ~/Downloads:/workdir vid-to-gif \
    sh -c "set -e && \
    echo 'Checking media type...' && \
    page_content=\$(curl -s -A 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36' '$input_url') && \
    if echo \"\$page_content\" | grep -q 'tweet_video_thumb'; then \
        # It's a native GIF
        thumb_url=\$(echo \"\$page_content\" | grep -o 'https://pbs.twimg.com/tweet_video_thumb/[^\"]*') && \
        gif_url=\${thumb_url/tweet_video_thumb/tweet_video} && \
        gif_url=\${gif_url/.jpg/.gif} && \
        filename=\$(basename \$gif_url) && \
        echo \"Downloading GIF from \$gif_url\" && \
        curl -L -o \"/workdir/\$filename\" \$gif_url && \
        echo \"Downloaded GIF: \$filename\"; \
    else \
        # It's a video, try alternative method
        echo 'Attempting to fetch video URL...' && \
        video_url=\$(yt-dlp -g '$input_url' --format 'best[ext=mp4]' 2>/dev/null) && \
        if [ -z \"\$video_url\" ]; then \
            echo \"Failed to get video URL, trying backup method...\" && \
            video_url=\$(yt-dlp -g '$input_url' --format 'bestvideo' 2>/dev/null); \
        fi && \
        if [ -z \"\$video_url\" ]; then exit 1; fi && \
        sleep 2 && \
        filename=\$(yt-dlp --get-filename -o '%(uploader)s_%(id)s' '$input_url') && \
        if [ -z \"\$filename\" ]; then filename=\"output\"; fi && \
        sleep 1 && \
        ffmpeg -i \"\$video_url\" \
        -vf 'fps=10,scale=480:-1:flags=lanczos' \
        -c:v gif \
        \"/workdir/\${filename}.gif\" && \
        echo \"Converted successfully to \${filename}.gif\"; \
    fi" || \
    { echo "Error: Failed to process media"; exit 1; }

if [ $? -eq 0 ]; then
    echo "Media saved to ~/Downloads/"
else
    echo "Processing failed"
    exit 1
fi