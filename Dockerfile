# Stage 1: Build stage - Compile the Go application
# Use official Go image as base (includes Go compiler and tools)
FROM golang:latest AS builder

# Set working directory inside container to /app
# All subsequent commands will run from this directory
WORKDIR /app

# Copy dependency files first (go.mod and go.sum)
# This is done before copying source code for better Docker layer caching
# If dependencies don't change, Docker can reuse cached layers
COPY go.mod go.sum ./

# Download all Go module dependencies
# This populates the module cache with required packages
RUN go mod download

# Copy all source code into the container
# First "." = source (your project root directory where you run "docker build")
# Second "." = destination (/app/ in container, set by WORKDIR above)
# This includes your application code (cmd/, internal/, etc.)
COPY . .

# Build the Go application
# CGO_ENABLED=0: Disable CGO for static binary (required for Alpine Linux)
# GOOS=linux: Build for Linux OS
# -o server: Output binary will be named "server"
# ./cmd/server: Build the package at ./cmd/server (contains main.go)
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server

# Stage 2: Runtime stage - Create minimal image with just the binary
# Use Alpine Linux (very small ~5MB base image)
FROM alpine:latest

# Install CA certificates (needed for HTTPS/TLS connections)
# --no-cache: Don't store package index locally (keeps image smaller)
RUN apk --no-cache add ca-certificates

# Set working directory for the final stage
WORKDIR /root/

# Copy the compiled binary from the builder stage
# --from=builder: Copy from the builder stage (first FROM statement)
# /app/server: Source file location in builder stage (created at line 27 with "go build -o server")
# .: Current directory (which is /root/), so binary will be at /root/server
COPY --from=builder /app/server .

# Copy Firebase service account credentials file
# This file contains credentials needed for Firebase authentication
# Make sure fb-sa-jobs-pk.json exists in your project root (even though it's in .gitignore)
COPY --from=builder /app/fb-sa-jobs-pk.json /root/fb-sa-jobs-pk.json

# Verify the binary exists and make it executable (just in case)
RUN chmod +x /root/server && ls -la /root/server

# Set the default command to run when container starts
# ["./server"]: Execute the server binary in the current directory
CMD ["./server"]