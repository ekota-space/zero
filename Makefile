DB_URL=postgres://postgres:postgres@localhost:5432/zero?sslmode=disable

diff:
		atlas migrate diff --env gorm

migrate:
		atlas migrate apply --env gorm --url $(DB_URL)

reset-db:
		atlas schema clean --url $(DB_URL)