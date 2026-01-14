#!/bin/bash

set -e

cd "$(dirname "$0")/../.."

OUTPUT_DIR="dist"
BINARY_NAME="app"

echo "🔨 Building Go binary..."

mkdir -p ${OUTPUT_DIR}

CGO_ENABLED=0 go build -ldflags="-s -w" -o ${OUTPUT_DIR}/${BINARY_NAME} cmd/app/main.go

echo "✅ Build completed successfully!"
echo "📦 Binary: ${OUTPUT_DIR}/${BINARY_NAME}"
echo ""
echo "Run with: ./${OUTPUT_DIR}/${BINARY_NAME}"
