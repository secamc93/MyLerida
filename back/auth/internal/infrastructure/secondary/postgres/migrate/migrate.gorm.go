package migrate

import (
	"auth/internal/infrastructure/secondary/postgres"
	"auth/internal/infrastructure/secondary/postgres/models"
	"auth/pkg/logger"
)

func Migrate() {
	log := logger.NewLogger()
	db := postgres.New().GetDB()
	err := db.AutoMigrate(
		&models.Users{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database: " + err.Error())
	}
	log.Info("Database migrated")
}
