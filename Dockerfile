# Use the Golang base image
FROM golang:1.20 AS build

# Install required dependencies
# Install required dependencies
RUN apt-get update && \
    apt-get install -y \
        libgl1-mesa-glx \
        libgl1-mesa-dri \
        libxxf86vm-dev \
        libgl1-mesa-dev \
        xorg-dev && \
    rm -rf /var/lib/apt/lists/*

# Set the working directory
WORKDIR /app

# Copy dependency files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the application code
COPY . .

# Build the application
RUN go build -o password-generator main.go

# Additional commands (if needed)
