FROM python:3.11-alpine

# Install ffmpeg and dependencies
RUN apk add --no-cache ffmpeg curl

# Install latest yt-dlp with custom settings
RUN pip install --no-cache-dir --upgrade yt-dlp && \
    mkdir -p /etc/yt-dlp && \
    echo "# yt-dlp configuration\n\
--sleep-interval 2\n\
--max-sleep-interval 5\n\
--user-agent 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36'\n\
--add-header 'Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8'\n\
--add-header 'Accept-Language: en-US,en;q=0.5'\n\
--socket-timeout 10" > /etc/yt-dlp/config

WORKDIR /workdir