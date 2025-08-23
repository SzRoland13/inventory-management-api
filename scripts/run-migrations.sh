#!/bin/sh

set -e

# Wait for Postgres to be ready
echo "Waiting for Postgres at $DB_HOST:$DB_PORT..."
for i in $(seq 1 30); do
    if pg_isready -h "$DB_HOST" -p "$DB_PORT" -U "$DB_USER"; then
        echo "Postgres is ready!"
        break
    fi
    echo "Attempt $i/30..."
    sleep 1
done

# Run migrations
echo "Running migrations..."
migrate -path ./db/migrations -database "postgres://$DB_USER:$DB_PASS@$DB_HOST:$DB_PORT/$DB_NAME?sslmode=disable" up
