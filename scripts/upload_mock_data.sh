#!/bin/bash

NAME="${CONTAINER_NAME:=bikesense-postgres}"
DB_USER="${POSTGRES_USER:=postgres}"
DB_PASSWORD="${POSTGRES_PASSWORD:=postgrespw}"
DB_NAME="${POSTGRES_DB:=bikesense}"
DB_PORT="${POSTGRES_PORT:=5432}"
DB_HOST="${POSTGRES_HOST:=localhost}"

export PGPASSWORD="${DB_PASSWORD}"
echo "Uploading mock data to the database..."
psql -U "${DB_USER}" -h "${DB_HOST}" -p "${DB_PORT}" -d "${DB_NAME}" -f mock_data.sql >/dev/null
echo "Mock data has been uploaded to the database!"
