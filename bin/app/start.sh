#!/bin/bash

set -e

cd "$(dirname "$0")/../.."

# Export variables solo si no están definidas (desarrollo local)
export DOCS_SERVICE_ADDR=${DOCS_SERVICE_ADDR:-localhost:8888}

echo "🚀 Starting markitos-it-app-website (Go)..."
echo "📡 DOCS_SERVICE_ADDR: $DOCS_SERVICE_ADDR"
echo ""

go run cmd/app/main.go
