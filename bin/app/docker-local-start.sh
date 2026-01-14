#!/bin/bash

set -e

cd "$(dirname "$0")/../.."

IMAGE_NAME="markitos-it-app-website"
TAG="local"
PORT=8080

echo "🚀 Starting Docker container: ${IMAGE_NAME}:${TAG}"
echo "📡 Port mapping: ${PORT}:${PORT}"
echo "🌐 Access at: http://localhost:${PORT}"

docker run --rm -p ${PORT}:${PORT} ${IMAGE_NAME}:${TAG}
