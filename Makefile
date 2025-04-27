.PHONY: start stop server-test migrate

MIGRATIONS_DIR=./server/internal/infrastructure/persistence/db/pgsql/migrations
NETWORK=wordle-turkish-overengineering_wordle-network
MIGRATE_IMAGE=migrate/migrate
DB_USER ?= postgres
DB_PASSWORD ?= postgres
DB_HOST ?= database
DB_PORT ?= 5432
DB_NAME ?= wordle

DB_URL := postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable

start:
	docker compose --env-file ./server/.env up --build
stop:
	docker compose --env-file ./server/.env down
server-test:
	docker compose --env-file ./server/.env exec server go test -v ./...
migrate:
	docker run --rm -v $(MIGRATIONS_DIR):/migrations --network $(NETWORK) \
		$(MIGRATE_IMAGE) -path=/migrations -database=$(DB_URL) up