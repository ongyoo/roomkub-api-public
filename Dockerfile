# Use the official Golang image to create a build artifact.
# This is based on Debian and sets the GOPATH to /go.
# https://hub.docker.com/_/golang
FROM golang:1.20.10 as builder

# Create and change to the app directory.
WORKDIR /app

# Retrieve application dependencies using go modules.
# Allows container builds to reuse downloaded dependencies.
COPY go.* ./
RUN go mod download

# Copy local code to the container image.
COPY . .

# Build the binary.
# -mod=readonly ensures immutable go.mod and go.sum in container builds.
# RUN CGO_ENABLED=0 GOOS=linux go build -mod=readonly -v -o server
# RUN CGO_ENABLED=0 go build -o webapp ./cmd/main.go
ENV GO111MODULE=on
RUN CGO_ENABLED=0 GOOS=linux go build -mod=readonly -v -o app ./cmd/main.go
#RUN CGO_ENABLED=0 go build -o webapp ./cmd/main.go

# Use the official Alpine image for a lean production container.
# https://hub.docker.com/_/alpine
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM alpine:3.12.0
RUN apk add --no-cache ca-certificates
WORKDIR /app

# Copy the binary to the production image from the builder stage.

COPY --from=builder /app .
RUN chmod a+x /app

# Expose port 8080 to the outside world
EXPOSE 3000

# Command to run the executable
CMD ["./app"]