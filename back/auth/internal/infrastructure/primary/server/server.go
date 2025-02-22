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

	dbConn := setupDatabase()
	defer closeDatabase(dbConn, log)

	migrate.Migrate()
	if port == "" {
		log.Fatal("Server port not set")
	}

	address := fmt.Sprintf(":%s", port)
	log.Success("Server running on port %s", port)
	log.Success("Swagger documentation available at http://localhost%s/swagger/index.html", address)

	setupSwagger(port)

	server := setupServer(address, log)
	startServer(server, log)
}

func setupDatabase() postgres.DBConnection {
	dbConn := postgres.New()
	return dbConn
}

func closeDatabase(dbConn postgres.DBConnection, log logger.ILogger) {
	if err := dbConn.CloseDB(); err != nil {
		log.Error("Error closing DB: %v", err)
	}
}

func setupSwagger(port string) {
	host := fmt.Sprintf("localhost:%s", port)
	docs.SwaggerInfo.Host = host
}

func setupServer(address string, log logger.ILogger) *http.Server {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = log.Writer()
	gin.DefaultErrorWriter = log.Writer()
	router := routers.SetupRouter()
	server := &http.Server{
		Addr:         address,
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	return server
}

func startServer(server *http.Server, log logger.ILogger) {
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
