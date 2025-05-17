.PHONY: up down server-test migrate migrate-down drop-db word-seed rotate-game up-detached cold-start

MIGRATIONS_DIR=./server/internal/infrastructure/persistence/db/pgsql/migrations
NETWORK=wordle-network
MIGRATE_IMAGE=migrate/migrate
DB_USER ?= postgres
DB_PASSWORD ?= postgres
DB_HOST ?= database
DB_PORT ?= 5432
DB_NAME ?= wordle

DB_URL := postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable

up:
	docker compose --env-file ./server/.env up --build
up-detached:
	docker compose --env-file ./server/.env up --build -d
down:
	docker compose --env-file ./server/.env down
server-test:
	docker compose --env-file ./server/.env exec server go test -v ./...
migrate:
	docker run --rm -v $(MIGRATIONS_DIR):/migrations --network $(NETWORK) \
		$(MIGRATE_IMAGE) -path=/migrations -database=$(DB_URL) up
migrate-down:
	docker run --rm -v $(MIGRATIONS_DIR):/migrations --network $(NETWORK) \
  $(MIGRATE_IMAGE) -path=/migrations -database=$(DB_URL) down -all
word-seed:
	docker run --rm \
	  --name task-runner \
	  --network $(NETWORK) \
	  -v ./server:/app \
	  --env-file ./server/.env \
	  -e ENV_FILE=/app/.env \
	  wordle-turkish-overengineering-server \
	  go run cmd/word/main.go
cold-start:
	bash ./cold_start.sh