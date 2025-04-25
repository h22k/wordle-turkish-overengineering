start:
	docker compose --env-file ./server/.env up --build
stop:
	docker compose --env-file ./server/.env down