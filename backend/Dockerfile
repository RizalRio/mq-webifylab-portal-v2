# ==========================================
# STAGE 1: Build the Go binary
# ==========================================
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copy dependency files
COPY go.mod go.sum ./
RUN go mod download

# Copy seluruh source code
COPY . .

# Build binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o server ./cmd/server/main.go

# ==========================================
# STAGE 2: Create minimal runtime image
# ==========================================
FROM alpine:latest

WORKDIR /app

# Install ca-certificates dan tzdata
RUN apk --no-cache add ca-certificates tzdata

# Copy binary
COPY --from=builder /app/server .

# Copy migrations folder (PENTING!)
COPY --from=builder /app/migrations ./migrations

# Expose port
EXPOSE 8080

# Jalankan binary
CMD ["./server"]