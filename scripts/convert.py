import sys
import subprocess
import os
import json
import base64
import re

def normalize_twitter_url(url):
    """Convert x.com URLs to twitter.com format"""
    return url.replace('x.com', 'twitter.com')

def convert_media(request_json):
    try:
        request = json.loads(request_json)
        url = request.get('url')
        mode = request.get('mode', 'gif')

        # Common options for yt-dlp
        yt_dlp_opts = [
            '--no-warnings',
            '--user-agent', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36',
            '--add-header', 'Accept:text/html,application/json,*/*',
            '--add-header', 'Accept-Language:en-US,en;q=0.9',
            '--no-playlist',
            '--format', 'best[ext=mp4]'
        ]

        # Get filename
        filename = subprocess.check_output(
            ['yt-dlp'] + yt_dlp_opts + ['--get-filename', '-o', '%(uploader)s_%(id)s', url],
            stderr=subprocess.PIPE
        ).decode().strip()

        if mode == 'gif':
            # For GIF mode, get video URL and convert
            video_url = subprocess.check_output(
                ['yt-dlp'] + yt_dlp_opts + ['-g', url],
                stderr=subprocess.PIPE
            ).decode().strip()

            process = subprocess.Popen([
                'ffmpeg', '-i', video_url,
                '-vf', 'fps=10,scale=480:-1:flags=lanczos',
                '-c:v', 'gif',
                '-f', 'gif',
                'pipe:1'
            ], stdout=subprocess.PIPE, stderr=subprocess.PIPE)
            
            media_data, err = process.communicate()
            output_ext = '.gif'
        else:
            # For video mode, download directly to memory
            process = subprocess.Popen(
                ['yt-dlp'] + yt_dlp_opts + ['-o', '-', url],
                stdout=subprocess.PIPE, stderr=subprocess.PIPE
            )
            media_data, err = process.communicate()
            output_ext = '.mp4'

        if process.returncode != 0:
            raise Exception(f"Process error: {err.decode()}")

        response = {
            'filename': f"{filename}{output_ext}",
            'data': base64.b64encode(media_data).decode('utf-8')
        }

        print(json.dumps(response))
        return True

    except Exception as e:
        print(json.dumps({'error': str(e)}), file=sys.stderr)
        return False

if __name__ == '__main__':
    # Read JSON request from stdin
    request_json = sys.stdin.read()
    success = convert_media(request_json)
    sys.exit(0 if success else 1) 