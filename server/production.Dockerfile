# 1. Build stage
FROM golang:1.23-alpine AS builder

# Install necessary packages
RUN apk add --no-cache git

WORKDIR /app

# Copy go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy application source code
COPY . .

# Build the binary (static, optimized)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o app ./cmd/api/main.go

# 2. Production stage
FROM scratch

# Copy CA certificates (needed for TLS/HTTPS)
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy the built binary from the builder stage
COPY --from=builder /app/app /app/app

# Run the binary
ENTRYPOINT ["/app/app"]