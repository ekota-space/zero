package db

import (
	"database/sql"
	"fmt"

	"github.com/ekota-space/zero/pkgs/common"
	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB

func SetupDatabaseConnection() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		common.Env.PostgresHost,
		common.Env.PostgresUser,
		common.Env.PostgresPassword,
		common.Env.PostgresDB,
		common.Env.PostgresPort,
	)

	db, err := sql.Open("pgx", dsn)

	if err != nil {
		panic(err)
	}

	DB = db
}
