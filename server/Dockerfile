FROM golang:1.23-alpine

RUN apk add --no-cache git curl && \
    go install github.com/air-verse/air@latest

WORKDIR /app
COPY . .

COPY .air.toml .

CMD ["air"]