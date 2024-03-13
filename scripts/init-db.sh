#!/bin/bash

# This script is used to initialize a postgres instance for development.
# For deployment, the docker-compose file should be used.

# enable debugging prints
set -x
# exit on error
set -eo pipefail

# Check for database variables otherwise use defaults
DB_USER="${POSTGRES_USER:=postgres}"
DB_PASSWORD="${POSTGRES_PASSWORD:=postgrespw}"
DB_NAME="${POSTGRES_DB:=bikesense}"
DB_PORT="${POSTGRES_PORT:=5432}"
DB_HOST="${POSTGRES_HOST:=localhost}"

if [[ -z "${SKIP_DOCKER}" ]]; then
	docker run \
		--name dev-postgres-esis \
		-e POSTGRES_USER=${DB_USER} \
		-e POSTGRES_PASSWORD=${DB_PASSWORD} \
		-e POSTGRES_DB=${DB_NAME} \
		-p "${DB_PORT}":5432 \
		-d postgres:16-alpine \
		postgres -N 1000
# ^ Increased maximum number of connections for testing purposes
fi

>&2 echo "Postgres has been initialized, ready to go!"
