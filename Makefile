.PHONY: start stop server-test migrate migrate-down drop-db word-seed rotate-game start-detached

MIGRATIONS_DIR=./server/internal/infrastructure/persistence/db/pgsql/migrations
NETWORK=wordle-network
MIGRATE_IMAGE=migrate/migrate
DB_USER ?= postgres
DB_PASSWORD ?= postgres
DB_HOST ?= database
DB_PORT ?= 5432
DB_NAME ?= wordle

DB_URL := postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable

start:
	docker compose --env-file ./server/.env up --build
start-detached:
	docker compose --env-file ./server/.env up --build -d
stop:
	docker compose --env-file ./server/.env down
server-test:
	docker compose --env-file ./server/.env exec server go test -v ./...
migrate:
	docker run --rm -v $(MIGRATIONS_DIR):/migrations --network $(NETWORK) \
		$(MIGRATE_IMAGE) -path=/migrations -database=$(DB_URL) up
migrate-down:
	docker run --rm -v $(MIGRATIONS_DIR):/migrations --network $(NETWORK) \
		$(MIGRATE_IMAGE) -path=/migrations -database=$(DB_URL) down -all
drop-db:
	docker run --rm --network $(NETWORK) \
		postgres:15-alpine psql -h $(DB_HOST) -U $(DB_USER) -c "DROP DATABASE IF EXISTS $(DB_NAME);" postgres && \
	docker run --rm --network $(NETWORK) \
		postgres:15-alpine psql -h $(DB_HOST) -U $(DB_USER) -c "CREATE DATABASE $(DB_NAME);" postgres
word-seed:
	docker run --rm \
	  --name task-runner \
	  --network $(NETWORK) \
	  -v ./server:/app \
	  --env-file ./server/.env \
	  wordle-turkish-overengineering-server \
	  go run cmd/word/main.go
