# Start from a Golang base image with version 1.22
FROM golang:1.22-alpine AS builder

# Set the working directory
WORKDIR /app

# Copy the Go module files
COPY go.mod ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY *.go ./

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Start a new stage from scratch
FROM ubuntu:latest  

# Install ca-certificates and bash for HTTPS and shell access
RUN apt-get update && apt-get install -y ca-certificates bash

WORKDIR /root/

# Copy the pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Set bash as the default shell when the container starts
ENTRYPOINT ["/bin/bash"]