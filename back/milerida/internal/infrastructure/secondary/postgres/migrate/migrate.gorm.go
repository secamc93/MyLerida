package migrate

import (
	"my-lerida/internal/infrastructure/secondary/postgres"
	"my-lerida/internal/infrastructure/secondary/postgres/models"
	"my-lerida/pkg/logger"
)

func Migrate() {
	log := logger.NewLogger()
	db := postgres.NewDBConnection().GetDB()
	err := db.AutoMigrate(
		&models.Restaurant{},
		&models.Service{},
		&models.Transportation{},
		&models.Tourism{})
	if err != nil {
		log.Fatal("Failed to migrate database: " + err.Error())
	}
	log.Success("Database migrated")
}
