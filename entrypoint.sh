#!/bin/sh

# Wait for the database to be ready
until nc -z -v -w30 postgres 5432
do
  echo "Waiting for the database connection..."
  sleep 1
done
echo "Database connection established"

# Run the migration
DATABASE_URL=postgres://user:password@postgres:5432/silver_bassoon?sslmode=disable ./bin/migrate

# Start the backend server
./bin/server
