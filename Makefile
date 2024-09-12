DB_URL=postgres://postgres:postgres@localhost:5432/zero?sslmode=disable
MIGRATION_DIR=file://db/migrations

gen:
		jet -dsn=$(DB_URL) -schema=public -path=./pkgs/root/db

diff:
		atlas migrate diff --env local

migrate:
		atlas migrate apply --env local --url "$(DB_URL)"

reset-db:
		atlas schema clean --url "$(DB_URL)"

status:
		atlas migrate status --env local --url="$(DB_URL)" --dir="$(MIGRATION_DIR)"

watch:
		gow run . --watch *.go