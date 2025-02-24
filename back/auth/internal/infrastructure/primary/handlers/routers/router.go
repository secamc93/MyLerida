package routers

import (
	"auth/internal/application/usecase/usecaseuser"
	userHandler "auth/internal/infrastructure/primary/handlers"
	"auth/internal/infrastructure/secondary/postgres"
	"auth/internal/infrastructure/secondary/postgres/repository"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Configurar CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

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
		api.POST("/login", handler.Login) // Nueva ruta de login
	}

	// Ruta para acceder a Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
