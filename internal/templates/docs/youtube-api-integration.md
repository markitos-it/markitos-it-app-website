# YouTube API Integration Guide

![YouTube API](https://images.unsplash.com/photo-1611162616305-c69b3fa7fbe0?w=1200&h=400&fit=crop)

## Overview

Learn how to integrate YouTube's Data API v3 into your applications to fetch videos, playlists, and channel information.

## Authentication Setup

### 1. Create API Credentials

Go to [Google Cloud Console](https://console.cloud.google.com) and:
- Create a new project
- Enable YouTube Data API v3
- Create credentials (API Key)

### 2. Install Dependencies

```bash
npm install googleapis
```

## Basic Implementation

```javascript
const { google } = require('googleapis');

const youtube = google.youtube({
  version: 'v3',
  auth: 'YOUR_API_KEY'
});

// Search for videos
async function searchVideos(query) {
  const response = await youtube.search.list({
    part: 'snippet',
    q: query,
    maxResults: 10,
    type: 'video'
  });
  
  return response.data.items;
}
```

## Fetching Video Details

```javascript
async function getVideoDetails(videoId) {
  const response = await youtube.videos.list({
    part: 'snippet,statistics,contentDetails',
    id: videoId
  });
  
  return response.data.items[0];
}
```

## Working with Playlists

```javascript
async function getPlaylistItems(playlistId) {
  const response = await youtube.playlistItems.list({
    part: 'snippet',
    playlistId: playlistId,
    maxResults: 50
  });
  
  return response.data.items;
}
```

## Rate Limits

YouTube API has quota limits:
- **10,000 units per day** (default)
- Search costs: 100 units
- Videos.list costs: 1 unit

## Best Practices

1. **Cache responses** to reduce API calls
2. **Use batch requests** when possible
3. **Handle errors gracefully** with retry logic
4. **Monitor your quota** usage

## Embedding Videos

```html
<iframe 
  width="560" 
  height="315" 
  src="https://www.youtube.com/embed/VIDEO_ID"
  frameborder="0" 
  allowfullscreen>
</iframe>
```

**Pro tip**: Use the `nocookie` domain for better privacy! 🎥
