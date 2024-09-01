# # Roomkub-api/cmd/roomkub-business-channel-api/Dockerfile

# # Use the official Golang image to create a build artifact
# FROM golang:1.20 as builder

# # Set the Current Working Directory inside the container
# WORKDIR /app

# # Copy go.mod and go.sum files
# COPY go.mod go.sum ./

# # Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
# RUN go mod download

# # Copy the source code
# COPY cmd/roomkub-business-channel-api/ .

# # Build the Go app
# RUN go build -o main .
# # ENV GO111MODULE=on
# # RUN CGO_ENABLED=0 GOOS=linux go build -mod=readonly -v -o app ./cmd/roomkub-business-channel-api/main.go

# # Start a new stage from scratch
# FROM alpine:latest

# # Set environment variables
# ENV ENV_PATH /.env

# # Copy the Pre-built binary file from the previous stage
# COPY --from=builder /app/main /app/main

# # Copy the .env file
# COPY .env /app/.env

# # Expose port 8080 to the outside world
# EXPOSE 8080

# # Command to run the executable
# CMD ["/app/main"]


# Start with a base Golang image
FROM golang:1.20.10 as builder

# Set necessary environment variables
# ENV GO111MODULE=on \
#     CGO_ENABLED=0 \
#     GOOS=linux \
#     GOARCH=amd64

ENV GO111MODULE=on \
    CGO_ENABLED=0

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to the working directory
COPY go.mod ./
COPY go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire source code to the working directory
COPY . .

# Build the Go app
RUN go build -mod=readonly -v -o app ./cmd/roomkub-business-channel-api/main.go

# Use a minimal base image to run the app
FROM alpine:3.15

# Set necessary environment variables
ENV GIN_MODE=release

# Copy the binary from the builder stage to the final stage
# COPY --from=builder /app/main /app/main

COPY --from=builder /app .
RUN chmod a+x /app

# Expose port 8080 to the outside world
EXPOSE 8081

# Command to run the executable
# CMD ["/app/main"]
CMD ["./app"]