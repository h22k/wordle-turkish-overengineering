FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copy only necessary files for building
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the game binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /game ./cmd/game/main.go

# Final stage
FROM alpine:3.19

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /game ./game
COPY --from=builder /app/.env ./.env

# Create log symlink to docker logs
RUN ln -sf /dev/stdout /var/log/cron.log

# Add crontab
RUN echo "* * * * * /app/game 2>&1" > /etc/crontabs/root

# Use busybox crond which is already included in alpine
USER root

# Start crond in foreground with more verbose logging
CMD ["/usr/sbin/crond", "-f", "-d", "8"] 