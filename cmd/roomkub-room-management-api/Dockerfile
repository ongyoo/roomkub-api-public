# Inherit from the base Dockerfile
FROM golang:1.20-alpine as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Set the service name
ARG SERVICE_NAME=roomkub-room-management-api

# Build the Go app
RUN go build -o /ROOMKUB-API cmd/roomkub-room-management-api/main.go

# Start a new stage from scratch
FROM alpine:latest

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /ROOMKUB-API ./

# Copy the .env file
COPY .env ./

# Export necessary environment variables
ENV SERVICE_NAME=roomkub-room-management-api

# Command to run the executable
CMD ["./roomkub-room-management-api"]
