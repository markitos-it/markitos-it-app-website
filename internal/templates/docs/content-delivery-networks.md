# Content Delivery Networks (CDN)

![CDN Network](https://images.unsplash.com/photo-1558494949-ef010cbdcc31?w=1200&h=400&fit=crop)

## What is a CDN?

A Content Delivery Network is a geographically distributed network of servers that deliver content to users based on their location, improving speed and reliability.

## How CDNs Work

```
User Request → DNS → Nearest Edge Server → Origin Server (if not cached)
    ↓              ↓           ↓                    ↓
Location      Route to    Check Cache         Fetch Content
Detection     Closest     Return if Hit       Cache & Return
```

## Benefits

### 1. Performance
- **Reduced latency**: Content served from nearby servers
- **Faster load times**: Cached static assets
- **Better UX**: Improved page speed

### 2. Reliability
- **High availability**: Redundant servers
- **DDoS protection**: Distributed traffic absorption
- **Failover**: Automatic rerouting

### 3. Cost Savings
- **Reduced bandwidth**: Less origin server traffic
- **Lower hosting costs**: Offload to CDN
- **Efficient scaling**: Handle traffic spikes

## Popular CDN Providers

### CloudFlare

```javascript
// Configure CloudFlare cache
// cloudflare-worker.js
addEventListener('fetch', event => {
  event.respondWith(handleRequest(event.request));
});

async function handleRequest(request) {
  const cache = caches.default;
  let response = await cache.match(request);
  
  if (!response) {
    response = await fetch(request);
    
    // Cache for 1 hour
    const headers = new Headers(response.headers);
    headers.set('Cache-Control', 'public, max-age=3600');
    
    response = new Response(response.body, {
      status: response.status,
      statusText: response.statusText,
      headers
    });
    
    event.waitUntil(cache.put(request, response.clone()));
  }
  
  return response;
}
```

### AWS CloudFront

```javascript
// CloudFront configuration
const cloudfront = require('aws-sdk/clients/cloudfront');

const params = {
  DistributionConfig: {
    CallerReference: Date.now().toString(),
    Comment: 'My CDN Distribution',
    DefaultCacheBehavior: {
      TargetOriginId: 'my-origin',
      ViewerProtocolPolicy: 'redirect-to-https',
      AllowedMethods: {
        Items: ['GET', 'HEAD', 'OPTIONS'],
        Quantity: 3
      },
      ForwardedValues: {
        QueryString: false,
        Cookies: { Forward: 'none' }
      },
      MinTTL: 0,
      DefaultTTL: 86400,
      MaxTTL: 31536000
    },
    Origins: {
      Quantity: 1,
      Items: [{
        Id: 'my-origin',
        DomainName: 'origin.example.com',
        CustomOriginConfig: {
          HTTPPort: 80,
          HTTPSPort: 443,
          OriginProtocolPolicy: 'https-only'
        }
      }]
    },
    Enabled: true
  }
};

cloudfront.createDistribution(params, (err, data) => {
  if (err) console.error(err);
  else console.log('Distribution created:', data.Distribution.Id);
});
```

## Cache Strategies

### 1. Cache-Control Headers

```javascript
// Express middleware
app.use((req, res, next) => {
  // Static assets - cache for 1 year
  if (req.path.match(/\.(jpg|jpeg|png|gif|css|js)$/)) {
    res.set('Cache-Control', 'public, max-age=31536000, immutable');
  }
  // HTML - revalidate
  else if (req.path.match(/\.html$/)) {
    res.set('Cache-Control', 'public, max-age=0, must-revalidate');
  }
  // API responses - short cache
  else if (req.path.startsWith('/api/')) {
    res.set('Cache-Control', 'public, max-age=300');
  }
  next();
});
```

### 2. Cache Invalidation

```javascript
// Purge CloudFlare cache
async function purgeCache(urls) {
  const response = await fetch(
    `https://api.cloudflare.com/client/v4/zones/${ZONE_ID}/purge_cache`,
    {
      method: 'POST',
      headers: {
        'X-Auth-Email': EMAIL,
        'X-Auth-Key': API_KEY,
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({ files: urls })
    }
  );
  
  return response.json();
}

// Usage
await purgeCache([
  'https://example.com/style.css',
  'https://example.com/script.js'
]);
```

### 3. Versioned Assets

```html
<!-- Version in filename -->
<link rel="stylesheet" href="/css/style.v2.3.1.css">
<script src="/js/app.v2.3.1.js"></script>

<!-- Or using query params -->
<link rel="stylesheet" href="/css/style.css?v=2.3.1">
<script src="/js/app.js?v=2.3.1"></script>

<!-- Best: Content hash -->
<link rel="stylesheet" href="/css/style.a3f89b.css">
<script src="/js/app.7d82c1.js"></script>
```

## Image Optimization

### Responsive Images

```html
<!-- Multiple resolutions -->
<img 
  srcset="
    image-320w.jpg 320w,
    image-640w.jpg 640w,
    image-1280w.jpg 1280w
  "
  sizes="(max-width: 320px) 280px,
         (max-width: 640px) 600px,
         1200px"
  src="image-640w.jpg"
  alt="Responsive image"
>
```

### Image CDN

```javascript
// Cloudinary transformation
const imageUrl = cloudinary.url('sample', {
  width: 800,
  height: 600,
  crop: 'fill',
  quality: 'auto',
  fetch_format: 'auto'
});
// Output: https://res.cloudinary.com/.../w_800,h_600,c_fill,q_auto,f_auto/sample

// imgix transformation
const imgixUrl = 'https://example.imgix.net/image.jpg' +
  '?w=800&h=600&fit=crop&auto=format,compress';
```

## Video Delivery

```javascript
// HLS video with CDN
const videoPlayer = videojs('my-video', {
  sources: [{
    src: 'https://cdn.example.com/video/playlist.m3u8',
    type: 'application/x-mpegURL'
  }],
  html5: {
    hls: {
      overrideNative: true,
      enableLowInitialPlaylist: true,
      smoothQualityChange: true
    }
  }
});
```

## Performance Monitoring

```javascript
// Measure CDN performance
async function checkCDNPerformance(url) {
  const start = performance.now();
  
  const response = await fetch(url, {
    cache: 'reload' // Bypass browser cache
  });
  
  const end = performance.now();
  const ttfb = end - start;
  
  const cdnCache = response.headers.get('cf-cache-status') || 
                   response.headers.get('x-cache');
  
  return {
    ttfb,
    cacheStatus: cdnCache,
    server: response.headers.get('server')
  };
}

// Usage
const metrics = await checkCDNPerformance('https://example.com/image.jpg');
console.log(`TTFB: ${metrics.ttfb}ms, Cache: ${metrics.cacheStatus}`);
```

## Security with CDN

### DDoS Protection

```javascript
// CloudFlare rate limiting
const rateLimitRule = {
  threshold: 100, // requests
  period: 60,     // seconds
  action: 'challenge', // or 'block'
  match: {
    request: {
      url: '/api/*'
    }
  }
};
```

### WAF Rules

```javascript
// Web Application Firewall
const wafRules = [
  {
    description: 'Block SQL injection',
    expression: '(http.request.uri.query contains "' + 
                'union select" or http.request.uri.query contains "1=1")',
    action: 'block'
  },
  {
    description: 'Block XSS attempts',
    expression: 'http.request.uri.query contains "<script"',
    action: 'block'
  }
];
```

## Best Practices

1. **Use long cache times** for versioned assets
2. **Implement cache warming** for critical content
3. **Monitor cache hit ratio** (aim for >90%)
4. **Compress assets** before CDN (gzip/brotli)
5. **Use HTTP/2** or HTTP/3 for multiplexing
6. **Enable TLS 1.3** for better security
7. **Set up monitoring** for CDN performance

**Deliver content at light speed! ⚡**
