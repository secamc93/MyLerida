package server

import (
	"fmt"
	"my-lerida/internal/infrastructure/primary/handlers/routers"
	"my-lerida/internal/infrastructure/secondary/postgres"
	"my-lerida/internal/infrastructure/secondary/postgres/migrate"
	"my-lerida/pkg/env"
	"my-lerida/pkg/logger"
	"net/http"
	"time"
)

func RunServer() {
	log := logger.NewLogger()
	dbConn := postgres.NewDBConnection()
	defer func() {
		if err := dbConn.CloseDB(); err != nil {
			log.Error("Error closing DB: %v", err)
		}
	}()
	migrate.Migrate()

	port := env.LoadEnv().ServerPort
	if port == "" {
		port = "8080"
	}

	address := fmt.Sprintf(":%s", port)
	log.Success("Server running on port %s", port)

	router := routers.SetupRouter()

	server := &http.Server{
		Addr:         address,
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Error("Failed to start server: %v", err)
	}
}
