# Stage 1: Build the Go binary
FROM golang:alpine AS builder

# Install ca-certificates (required for making HTTPS calls, e.g. to Telegram & Resend)
RUN apk add --no-cache ca-certificates

WORKDIR /app

# Copy go.mod and go.sum to cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go binary
# CGO_ENABLED=0 ensures a static binary, crucial for running in a minimal Alpine/Scratch image
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o webifylab-api ./cmd/server/main.go

# Stage 2: Create a minimal runner image
FROM alpine:latest

WORKDIR /app

# Copy CA certificates from builder stage
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy the compiled binary from builder stage
COPY --from=builder /app/webifylab-api .

# Expose port (opsional, sebagai dokumentasi)
EXPOSE 8080

# Start the application
CMD ["./webifylab-api"]
