include .env
export

PG_URL="postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=${POSTGRES_SSLMODE}"

env:
	@env

migrate-up:
	@migrate -database ${PG_URL} -path db/migrations up
migrate-down:
	@migrate -database ${PG_URL} -path db/migrations down

test:
	@go test ./...

dev:
	@go run ./cmd/main.go