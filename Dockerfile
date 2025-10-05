# Compilation environment for systems without Go
FROM golang:1.21

# Install necessary packages for OpenGL and GLFW compilation
RUN apt-get update && apt-get install -y \
    gcc \
    libx11-dev \
    libxrandr-dev \
    libxinerama-dev \
    libxcursor-dev \
    libxi-dev \
    libgl1-mesa-dev \
    xorg-dev \
    && rm -rf /var/lib/apt/lists/*

# Enable Go build cache mount
ENV GOCACHE=/tmp/go-cache

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum for dependency caching
COPY go.mod go.sum ./

# Download dependencies (cached if go.mod/go.sum unchanged)
RUN go mod download

# Copy source code
COPY . .

# Build application with optimizations
RUN --mount=type=cache,target=/tmp/go-cache \
    CGO_ENABLED=1 GOOS=linux go build \
    -ldflags="-s -w" \
    -o humangl ./cmd/humangl

# The binary will be copied out by the Makefile