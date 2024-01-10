# Stage 1: Build the application
FROM golang:1.21 AS builder

WORKDIR /app

# Copy the go.mod and go.sum files to download dependencies
COPY go.mod go.sum ./

RUN go mod download

# Copy the source code to the container
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o app .

# Stage 2: Create a minimal image to run the application
FROM alpine:latest

WORKDIR /app

# Copy the built binary from the previous stage
COPY --from=builder /app/app .

# Set the entrypoint to run the application
ENTRYPOINT ["./app"]
