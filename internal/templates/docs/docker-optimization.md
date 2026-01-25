# Docker Image Optimization

![Docker](https://images.unsplash.com/photo-1605745341112-85968b19335b?w=1200&h=400&fit=crop)

## Why Optimize Docker Images?

Smaller images mean:
- **Faster deployments**
- **Reduced storage costs**
- **Better security** (smaller attack surface)
- **Quicker CI/CD pipelines**

## Multi-Stage Builds

The most powerful optimization technique:

```dockerfile
# Build stage
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o myapp

# Production stage
FROM alpine:3.19
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/myapp .
CMD ["./myapp"]
```

## Choose the Right Base Image

```dockerfile
# ❌ Bad: Full OS (1.2GB)
FROM ubuntu:22.04

# ✅ Better: Slim variant (200MB)
FROM node:20-slim

# ✅ Best: Alpine (50MB)
FROM node:20-alpine
```

## Layer Optimization

Order matters! Put frequently changing layers last:

```dockerfile
FROM node:20-alpine

# 1. Install system dependencies (rarely changes)
RUN apk add --no-cache python3 make g++

# 2. Copy package files (changes occasionally)
COPY package*.json ./

# 3. Install dependencies
RUN npm ci --only=production

# 4. Copy application code (changes frequently)
COPY . .

CMD ["node", "index.js"]
```

## .dockerignore File

Exclude unnecessary files:

```
node_modules
npm-debug.log
.git
.gitignore
README.md
.env
.DS_Store
*.md
```

## Remove Build Dependencies

```dockerfile
FROM alpine:3.19

# Install, use, and remove in one layer
RUN apk add --no-cache --virtual .build-deps \
    gcc \
    musl-dev \
    && apk add --no-cache libffi \
    && pip install mypackage \
    && apk del .build-deps
```

## Use Specific Tags

```dockerfile
# ❌ Never do this
FROM node:latest

# ✅ Always pin versions
FROM node:20.11.0-alpine3.19
```

## Minimize Layers

Combine commands when possible:

```dockerfile
# ❌ Multiple layers
RUN apt-get update
RUN apt-get install -y curl
RUN apt-get install -y git

# ✅ Single layer
RUN apt-get update && \
    apt-get install -y curl git && \
    rm -rf /var/lib/apt/lists/*
```

## Security Scanning

```bash
# Scan with Trivy
trivy image myapp:latest

# Scan with Docker Scout
docker scout cves myapp:latest
```

## Benchmarking Results

| Technique | Before | After | Savings |
|-----------|--------|-------|---------|
| Multi-stage build | 1.2GB | 15MB | 98.7% |
| Alpine base | 500MB | 80MB | 84% |
| Layer optimization | 300MB | 250MB | 16.6% |

**Remember: Every MB saved is a deployment accelerated! 🚀**
