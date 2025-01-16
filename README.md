# Twitter-to-GIF

Convert Twitter/X.com videos to GIFs using Docker.

## The Mission
...is to create a browser plugin and mobile app that allows users to consume public social media posts, particularly memes, without the distractions of pop-ups, ads, and tracking. We aim to provide a clean, user-friendly experience that aggregates content from various platforms while prioritizing user privacy and customization.

## To do
- Add support for videos from various platforms:
  - Instagram
  - TikTok
  - YouTube
  - Twitch
  - Facebook
  - Reddit
  - LinkedIn
  - Vimeo
  - Dailymotion

### Media Management
- Ensure experience is ad-free
  - Download or stream from the Docker container
  - Consider AI-based solutions for in-video ad removal (e.g., using TensorFlow for ad-blocking within videos)
    - Imagine watching delayed sports but it has no ads. That's the dream.
- Enable offline access for downloaded content
- Facilitate user-generated content uploads

### Content Aggregation
- Create a system for aggregating memes
  - Develop customizable feed options
  - Implement search functionality for memes
  - Allow bookmarking and sharing of posts
  - Provide notifications for trending memes

### Privacy and Security
- Ensure privacy-focused design and functionality

### Cross-Platform Sync
- Ensure cross-platform sync for mobile app

### User Experience
- Conduct user testing and feedback sessions

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
