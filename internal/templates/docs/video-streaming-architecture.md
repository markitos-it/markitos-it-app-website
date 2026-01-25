# Video Streaming Architecture

![Video Streaming](https://images.unsplash.com/photo-1574717024653-61fd2cf4d44d?w=1200&h=400&fit=crop)

## Overview

Building a scalable video streaming platform requires careful architecture decisions. Let's explore how platforms like YouTube and Netflix handle billions of video plays.

## Architecture Components

### 1. Video Upload & Processing

```
User Upload → CDN → Processing Pipeline → Storage
```

**Processing steps:**
- Transcoding (multiple resolutions)
- Thumbnail generation
- Metadata extraction
- Quality analysis

### 2. Transcoding Pipeline

```javascript
const ffmpeg = require('fluent-ffmpeg');

function transcodeVideo(inputPath, outputPath) {
  return new Promise((resolve, reject) => {
    ffmpeg(inputPath)
      .output(`${outputPath}/720p.mp4`)
      .videoCodec('libx264')
      .size('1280x720')
      .videoBitrate('2500k')
      .output(`${outputPath}/480p.mp4`)
      .size('854x480')
      .videoBitrate('1000k')
      .output(`${outputPath}/360p.mp4`)
      .size('640x360')
      .videoBitrate('600k')
      .on('end', resolve)
      .on('error', reject)
      .run();
  });
}
```

## Adaptive Bitrate Streaming (ABR)

Use HLS or DASH protocols:

```bash
# Generate HLS playlist
ffmpeg -i input.mp4 \
  -c:v libx264 -c:a aac \
  -f hls \
  -hls_time 10 \
  -hls_playlist_type vod \
  -hls_segment_filename "segment%d.ts" \
  playlist.m3u8
```

### HLS Playlist Example

```
#EXTM3U
#EXT-X-VERSION:3
#EXT-X-TARGETDURATION:10
#EXTINF:10.0,
segment0.ts
#EXTINF:10.0,
segment1.ts
#EXTINF:10.0,
segment2.ts
#EXT-X-ENDLIST
```

## CDN Strategy

Distribute content globally:

```
User → Edge Server → Origin Server
  ↓        ↓              ↓
Cache   Cache Hot      Cold Storage
Layer   Content        (S3, GCS)
```

**Popular CDNs:**
- CloudFlare
- AWS CloudFront
- Fastly
- Akamai

## Database Schema

```sql
-- Videos table
CREATE TABLE videos (
  id UUID PRIMARY KEY,
  title VARCHAR(255),
  description TEXT,
  duration INTEGER,
  uploader_id UUID,
  upload_date TIMESTAMP,
  view_count BIGINT DEFAULT 0,
  status VARCHAR(50)
);

-- Video quality variants
CREATE TABLE video_variants (
  id UUID PRIMARY KEY,
  video_id UUID REFERENCES videos(id),
  quality VARCHAR(10), -- '1080p', '720p', etc.
  url TEXT,
  size_bytes BIGINT,
  codec VARCHAR(50)
);

-- View analytics
CREATE TABLE video_views (
  id UUID PRIMARY KEY,
  video_id UUID REFERENCES videos(id),
  user_id UUID,
  watched_duration INTEGER,
  quality_selected VARCHAR(10),
  timestamp TIMESTAMP
);
```

## Player Implementation

```javascript
// Video.js player setup
const player = videojs('my-video', {
  controls: true,
  autoplay: false,
  preload: 'auto',
  fluid: true,
  plugins: {
    qualitySelector: {
      default: '720p'
    }
  }
});

// Track analytics
player.on('play', () => {
  analytics.track('video_play', {
    videoId: currentVideoId,
    quality: player.currentQuality()
  });
});

player.on('timeupdate', () => {
  const progress = (player.currentTime() / player.duration()) * 100;
  if (progress > 25 && !milestones.quarter) {
    analytics.track('video_25_percent', { videoId: currentVideoId });
    milestones.quarter = true;
  }
});
```

## Scalability Considerations

### Storage
- Use object storage (S3, GCS) for videos
- Implement lifecycle policies for old content
- Consider archive storage for rarely accessed videos

### Processing
- Use worker queues (RabbitMQ, AWS SQS)
- Parallel processing for different qualities
- Retry mechanism for failed jobs

### Caching Strategy
```
L1: Browser cache (5 min)
L2: CDN edge (1 hour)
L3: CDN origin (24 hours)
L4: Object storage (permanent)
```

## Cost Optimization

1. **Lazy transcoding**: Only create qualities when requested
2. **Compression**: Use modern codecs (H.265, AV1)
3. **Smart CDN**: Route based on cost and performance
4. **Storage tiers**: Move old content to cheaper storage

## Security

- **DRM** for premium content
- **Token-based URLs** with expiration
- **HTTPS** everywhere
- **Rate limiting** on API endpoints

**Streaming millions of videos daily! 📺**
