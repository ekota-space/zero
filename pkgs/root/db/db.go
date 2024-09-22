package db

import (
	"database/sql"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/ekota-space/zero/pkgs/common"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func SetupDatabaseConnection() *sql.DB {
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

	return db
}

func RunMigrations() {
	db := SetupDatabaseConnection()

	defer db.Close()

	dbUri := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		common.Env.PostgresUser,
		common.Env.PostgresPassword,
		common.Env.PostgresHost,
		common.Env.PostgresPort,
		common.Env.PostgresDB,
	)

	_, err := db.Exec("DROP SCHEMA IF EXISTS public CASCADE")

	if err != nil {
		panic(err)
	}
	_, err = db.Exec("DROP SCHEMA IF EXISTS atlas_schema_revisions CASCADE")

	if err != nil {
		panic(err)
	}

	_, err = db.Exec("CREATE SCHEMA IF NOT EXISTS public")

	if err != nil {
		panic(err)
	}

	cmd := exec.Command("atlas", "migrate", "apply", "--env", "local", "--url", dbUri)
	cwd, err := os.Getwd()

	if err != nil {
		panic(err)
	}
	cmd.Dir = strings.Split(cwd, "/tests")[0]
	stdout, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println(string(stdout))
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
