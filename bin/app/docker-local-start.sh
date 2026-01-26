#!/bin/bash

set -e

cd "$(dirname "$0")/../.."

IMAGE_NAME="markitos-it-app-website"
TAG="local"
PORT=8080

# Export variables solo si no están definidas (desarrollo local)
export DOCS_SERVICE_ADDR=${DOCS_SERVICE_ADDR:-host.docker.internal:8888}

echo "🚀 Starting Docker container: ${IMAGE_NAME}:${TAG}"
echo "📡 Port mapping: ${PORT}:${PORT}"
echo "📡 DOCS_SERVICE_ADDR: $DOCS_SERVICE_ADDR"
echo "🌐 Access at: http://localhost:${PORT}"
echo ""

docker run --rm \
  -p ${PORT}:${PORT} \
  -e DOCS_SERVICE_ADDR=${DOCS_SERVICE_ADDR} \
  ${IMAGE_NAME}:${TAG}
