package root

import (
	"fmt"

	auth "github.com/ekota-space/zero/pkgs/auth/models"
	"github.com/ekota-space/zero/pkgs/common"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupDatabaseConnection() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		common.Env.PostgresHost,
		common.Env.PostgresUser,
		common.Env.PostgresPassword,
		common.Env.PostgresDB,
		common.Env.PostgresPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}

	DB = db

	RunMigrations()
}

func RunMigrations() {
	err := DB.AutoMigrate(&auth.Users{})

	if err != nil {
		panic(fmt.Sprintf("Failed to run migrations: %v", err))
	}
}
