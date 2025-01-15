# Twitter-to-GIF

Convert Twitter/X.com videos to GIFs using Docker.

## Prerequisites
- Docker
- Git

## Quick Start
  ```bash
  # Clone the repo
  git clone git@github.com:anubiskhan/twitter-to-gif.git
  cd twitter-to-gif

  # Make script executable
  chmod +x vid-to-gif.sh

  # Run (first run will build Docker image)
  ./vid-to-gif.sh https://x.com/hperryhorton/status/1747705965666079069
  ```

GIFs are saved to `~/Downloads/`.

## Notes
- Works with public Twitter/X.com video/gif posts
- Output format: `username_tweetid.gif`
- Default size: 480px width (maintains aspect ratio)
- FPS: 10