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

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum for dependency caching
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build application
RUN CGO_ENABLED=1 GOOS=linux go build -o humangl ./cmd/humangl

# The binary will be copied out by the Makefile