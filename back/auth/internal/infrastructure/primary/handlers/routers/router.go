package routers

import (
	"auth/internal/application/usecase/usecaseuser"
	userHandler "auth/internal/infrastructure/primary/handlers"
	"auth/internal/infrastructure/secondary/postgres"
	"auth/internal/infrastructure/secondary/postgres/repository"
	"os"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	dbConnection := postgres.New()
	userRepo := repository.New(dbConnection)
	userUseCase := usecaseuser.New(userRepo)

	urlBase := os.Getenv("URL_BASE")
	if urlBase == "" {
		urlBase = "/api"
	}

	api := router.Group(urlBase)
	{
		handler := userHandler.New(userUseCase)
		api.POST("/users", handler.CreateUser)
	}

	// Ruta para acceder a Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
