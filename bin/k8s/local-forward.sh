#!/bin/bash

set -e

echo "🔌 Forwarding K8s service to localhost:8080..."
echo "   Access the app at: http://localhost:8080"
echo "   Press Ctrl+C to stop"
echo ""

kubectl port-forward svc/markitos-it-app-website-service 8080:8080
