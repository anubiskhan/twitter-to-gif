import sys
import subprocess
import os
import json
import base64
import re

def normalize_twitter_url(url):
    """Convert x.com URLs to twitter.com format"""
    return url.replace('x.com', 'twitter.com')

def convert_media(url):
    try:
        # Normalize URL
        url = normalize_twitter_url(url)

        # Common options for yt-dlp
        yt_dlp_opts = [
            '--no-warnings',
            '--user-agent', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36',
            '--add-header', 'Accept:text/html,application/json,*/*',
            '--add-header', 'Accept-Language:en-US,en;q=0.9',
            '--add-header', 'Origin:https://twitter.com',
            '--add-header', 'Referer:https://twitter.com/'
        ]

        # Get filename from yt-dlp
        filename = subprocess.check_output(
            ['yt-dlp'] + yt_dlp_opts + ['--get-filename', '-o', '%(uploader)s_%(id)s', url],
            stderr=subprocess.PIPE
        ).decode().strip()

        # Get video URL
        video_url = subprocess.check_output(
            ['yt-dlp'] + yt_dlp_opts + ['-g', url, '--format', 'best[ext=mp4]'],
            stderr=subprocess.PIPE
        ).decode().strip()

        # Convert to GIF and output to memory/pipe
        ffmpeg_process = subprocess.Popen([
            'ffmpeg', '-i', video_url,
            '-vf', 'fps=10,scale=480:-1:flags=lanczos',
            '-c:v', 'gif',
            '-f', 'gif',
            'pipe:1'  # Output to stdout
        ], stdout=subprocess.PIPE, stderr=subprocess.PIPE)

        gif_data, err = ffmpeg_process.communicate()

        if ffmpeg_process.returncode != 0:
            raise Exception(f"FFmpeg error: {err.decode()}")

        # Create response with metadata and base64 encoded GIF
        response = {
            'filename': f"{filename}.gif",
            'data': base64.b64encode(gif_data).decode('utf-8')
        }

        # Output JSON response
        print(json.dumps(response))
        return True
    except subprocess.CalledProcessError as e:
        error_msg = e.stderr.decode() if e.stderr else str(e)
        print(json.dumps({'error': f"yt-dlp error: {error_msg}"}), file=sys.stderr)
        return False
    except Exception as e:
        print(json.dumps({'error': str(e)}), file=sys.stderr)
        return False

if __name__ == '__main__':
    if len(sys.argv) != 2:
        print("Usage: convert.py <url>")
        sys.exit(1)

    success = convert_media(sys.argv[1])
    sys.exit(0 if success else 1) 