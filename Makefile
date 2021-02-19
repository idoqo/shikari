include .env
export

PG_URL="postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${MIGRATE_POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=${POSTGRES_SSLMODE}"

env:
	@env

migrate-up:
	@migrate -database ${PG_URL} -path db/migrations up
migrate-down:
	@migrate -database ${PG_URL} -path db/migrations down
seed:
	@PGPASSWORD=${POSTGRES_PASSWORD} psql -h ${MIGRATE_POSTGRES_HOST} -U ${POSTGRES_USER} -d ${POSTGRES_DB} -f db/seeds/000001_seed_tables.up.sql
unseed:
	@PGPASSWORD=${POSTGRES_PASSWORD} psql -h ${MIGRATE_POSTGRES_HOST} -U ${POSTGRES_USER} -d ${POSTGRES_DB} -f db/seeds/000001_seed_tables.down.sql

test:
	@go test ./...

dev:
	@go build -o build/shikari ./cmd/ && ./build/shikari