#!/bin/bash
set -e

DB_CONTAINER_NAME="inventory_db"
DB_USER="postgres"
DB_PASS="postgres"
DB_NAME="inventory"
DB_HOST="localhost"
MIGRATIONS_DIR="./db/migrations"

echo "Starting database..."
docker compose up -d db

echo "Waiting for Postgres to be ready..."
until docker exec $DB_CONTAINER_NAME pg_isready -U $DB_USER > /dev/null 2>&1; do
    sleep 1
done

echo "Postgres is ready."

docker exec -e PGPASSWORD=$DB_PASS $DB_CONTAINER_NAME psql -U $DB_USER -tc "SELECT 1 FROM pg_database WHERE datname = '$DB_NAME'" | grep -q 1 || \
    docker exec -e PGPASSWORD=$DB_PASS $DB_CONTAINER_NAME psql -U $DB_USER -c "CREATE DATABASE $DB_NAME"

echo "Running migrations..."
docker compose run -T --rm migrate

echo "Setup complete."