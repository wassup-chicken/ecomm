#!/bin/bash
# Build script for Elastic Beanstalk deployment

echo "Building Go application for Linux..."
mkdir -p bin
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/application ./cmd/server
chmod +x bin/application
echo "Build complete! Binary created at bin/application"
