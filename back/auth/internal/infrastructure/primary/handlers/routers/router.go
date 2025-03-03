package routers

import (
	"auth/internal/application/usecase/usecaserole"
	"auth/internal/application/usecase/usecaseuser"

	"auth/internal/infrastructure/primary/handlers/handlers/rolehandlers"
	"auth/internal/infrastructure/primary/handlers/handlers/userhandlers/handlers"
	"auth/internal/infrastructure/primary/handlers/middleware"
	"auth/internal/infrastructure/secondary/postgres/connectpostgres"
	"auth/internal/infrastructure/secondary/postgres/repository"
	"auth/pkg/logger"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	log := logger.NewLogger()
	router := gin.New()
	router.Use(gin.RecoveryWithWriter(log.Writer()))
	router.Use(middleware.GinLogger(log))

	// Configurar CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	dbConnection := connectpostgres.New()

	repo := repository.New(dbConnection)
	userRepo := repo.NewUserRepository()
	roleRepo := repo.NewRoleRepository()

	userUseCase := usecaseuser.New(userRepo)
	roleUseCase := usecaserole.New(roleRepo)

	urlBase := os.Getenv("URL_BASE")
	if urlBase == "" {
		urlBase = "/api"
	}

	api := router.Group(urlBase)
	{
		// Rutas de usuarios
		handler := handlers.New(userUseCase)
		api.POST("/users", handler.CreateUser)
		api.GET("/users", handler.GetUsers)
		api.GET("/users/:id", handler.GetUser)
		api.PUT("/users/:id", handler.UpdateUser)
		api.DELETE("/users/:id", handler.DeleteUser)
		api.POST("/login", handler.Login)

		// Registrar rutas de roles
		rolehandlers.RegisterRoutes(api, roleUseCase)
	}

	// Configuración de Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
