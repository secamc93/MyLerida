package migrate

import (
	"auth/internal/infrastructure/secondary/postgres/connectpostgres"
	"auth/internal/infrastructure/secondary/postgres/models"
	"auth/pkg/logger"
)

func Migrate() {
	log := logger.NewLogger()
	db := connectpostgres.New().GetDB()
	err := db.AutoMigrate(
		&models.Module{},
		&models.Permission{},
		&models.Role{},
		&models.Users{},
		&models.Business{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database: " + err.Error())
	}
	log.Info("Database migrated")
}
