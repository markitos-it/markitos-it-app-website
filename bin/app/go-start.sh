#!/bin/bash

set -e

cd "$(dirname "$0")/../.."

echo "🚀 Starting markitos-it-app-website (Go)..."
go run cmd/app/main.go
