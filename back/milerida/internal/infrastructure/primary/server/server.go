package server

import (
	"context"
	"fmt"
	"my-lerida/internal/infrastructure/primary/handlers/routers"
	"my-lerida/internal/infrastructure/secondary/postgres"
	"my-lerida/internal/infrastructure/secondary/postgres/migrate"
	"my-lerida/pkg/env"
	"my-lerida/pkg/logger"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
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
		log.Fatal("Server port not set")
	}

	address := fmt.Sprintf(":%s", port)
	log.Success("Server running on port %s", port)

	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = log.Writer()
	router := routers.SetupRouter()
	server := &http.Server{
		Addr:         address,
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Failed to start server: %v", err)
		}
	}()

	<-quit
	log.Info("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: %v", err)
	}

	log.Success("Server exiting")
}
