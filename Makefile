DB_USER=admin
DB_PASS=admin123%21%40%23
DB_HOST=host.docker.internal
DB_PORT=5432
DB_NAME=chatapp
SSL_MODE=disable

DB_URL=postgres://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(SSL_MODE)
MIGRATE=docker run --rm -v $(PWD)/db/migrations:/migrations \
          migrate/migrate -path=/migrations -database "$(DB_URL)"

GENERATE_SOURCE_SQL=docker run --rm -v "$$PWD":/src -w /src sqlc/sqlc generate

sqlc:
	$(GENERATE_SOURCE_SQL)

migrate-up:
	$(MIGRATE) up

migrate-down:
	$(MIGRATE) down 1