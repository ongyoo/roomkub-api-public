# # Roomkub-api/cmd/roomkub-user-and-customer-api/Dockerfile
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
RUN go build -mod=readonly -v -o app ./cmd/roomkub-user-and-customer-api/main.go

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