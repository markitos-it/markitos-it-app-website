#!/bin/bash

set -e

cd "$(dirname "$0")/../.."

IMAGE_NAME="markitos-it-app-website"
TAG="local"
DIST_DIR="dist"

echo "🧹 Cleaning up..."

if [ -d "${DIST_DIR}" ]; then
    echo "  ➜ Removing ${DIST_DIR}/ directory..."
    rm -rf ${DIST_DIR}
fi

if docker image inspect ${IMAGE_NAME}:${TAG} >/dev/null 2>&1; then
    echo "  ➜ Removing Docker image ${IMAGE_NAME}:${TAG}..."
    docker rmi ${IMAGE_NAME}:${TAG}
fi

echo "✅ Cleanup completed!"
