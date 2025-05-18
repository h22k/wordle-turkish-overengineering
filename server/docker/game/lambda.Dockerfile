# 1. Build stage
FROM golang:1.23-alpine AS builder

# Install necessary packages
RUN apk add --no-cache git tzdata

WORKDIR /app

# Copy go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy application source code
COPY . .

# Build the binary (static, optimized)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o app ./cmd/lambda/main.go

# 2. Production stage
FROM scratch

# Copy timezone data (if needed)
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

# Copy the built binary from the builder stage
COPY --from=builder /app/app /app/app

# Run the binary
ENTRYPOINT ["/app/app"]