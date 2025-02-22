package server

import (
	"auth/internal/infrastructure/primary/handlers/routers"
	"auth/internal/infrastructure/secondary/postgres"
	"auth/internal/infrastructure/secondary/postgres/migrate"
	"auth/pkg/docs"
	"auth/pkg/env"
	"auth/pkg/logger"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

func RunServer() {
	log := logger.NewLogger()
	config := env.LoadEnv()
	port := config.ServerPort

	dbConn := postgres.New()
	defer func() {
		if err := dbConn.CloseDB(); err != nil {
			log.Error("Error closing DB: %v", err)
		}
	}()
	migrate.Migrate()
	if port == "" {
		log.Fatal("Server port not set")
	}

	address := fmt.Sprintf(":%s", port)
	log.Success("Server running on port %s", port)
	log.Info("Swagger documentation available at http://localhost%s/swagger/index.html", address)

	// Configurar dinámicamente el host para Swagger
	host := fmt.Sprintf("localhost:%s", port)
	docs.SwaggerInfo.Host = host

	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = log.Writer()
	gin.DefaultErrorWriter = log.Writer() // Asegurarse de que los errores también se registren
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
