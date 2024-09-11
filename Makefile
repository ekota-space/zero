DB_URL=postgres://postgres:postgres@localhost:5432/zero?sslmode=disable
MIGRATION_DIR=file://db/migrations

diff:
		atlas migrate diff --env gorm

migrate:
		atlas migrate apply --env gorm --url "$(DB_URL)"

reset-db:
		atlas schema clean --url "$(DB_URL)"

status:
		atlas migrate status --env gorm --url="$(DB_URL)" --dir="$(MIGRATION_DIR)"

watch:
		gow run . --watch *.go