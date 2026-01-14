#!/bin/bash

set -e

cd "$(dirname "$0")/../.."

IMAGE_NAME="markitos-it-app-website"
TAG="local"

echo "🏗️  Building Docker image: ${IMAGE_NAME}:${TAG}..."

docker build -t ${IMAGE_NAME}:${TAG} .

echo "✅ Build completed successfully!"
echo "📦 Image: ${IMAGE_NAME}:${TAG}"
