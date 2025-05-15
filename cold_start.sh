#!/bin/bash

set -e

echo "🔍 Checking .env file..."
if [ ! -f ./server/.env ]; then
  echo "⚠️  .env not found, copying from .env.example..."
  cp ./server/.env.example ./server/.env
fi

echo "🔍 Checking if Docker Compose is running..."
if [ -z "$(docker compose --env-file ./server/.env ps -q database)" ]; then
  echo "🚀 Starting Docker Compose..."
  make up-detached
else
  echo "✅ Docker Compose already running."
fi

echo "🗃️ MIGRATING database..."
make migrate
echo "✅ Migration done."

echo "🌱 Seeding database..."
make word-seed
echo "✅ Seeding done."

echo "🔁 First game initialization..."
if [ ! -f .game_rotator_initialized ]; then
  echo "🎮 First-time initialization of game-rotator..."
  docker compose --env-file ./server/.env exec -it game-rotator /app/game
  touch .game_rotator_initialized
else
  echo "✅ game-rotator already initialized."
fi
