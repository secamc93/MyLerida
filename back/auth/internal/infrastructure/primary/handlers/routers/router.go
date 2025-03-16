package routers

import (
	"auth/internal/application/usecase/usecasepermissions"
	"auth/internal/application/usecase/usecaserole"
	"auth/internal/application/usecase/usecaseuser"

	"auth/internal/infrastructure/primary/handlers/handlers/permissionshandlers/permissionsroutes"
	"auth/internal/infrastructure/primary/handlers/handlers/rolehandlers/roleroutes"
	"auth/internal/infrastructure/primary/handlers/handlers/userhandlers/userroutes"
	"auth/internal/infrastructure/primary/handlers/middleware"
	"auth/internal/infrastructure/secondary/postgres/connectpostgres"
	"auth/internal/infrastructure/secondary/postgres/repository"
	"auth/pkg/logger"

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
	permissionRepo := repo.NewPermissionRepository()

	userUseCase := usecaseuser.New(userRepo, permissionRepo)
	roleUseCase := usecaserole.New(roleRepo)
	permissionsUseCase := usecasepermissions.New(permissionRepo)

	api := router.Group("/api")
	{
		userroutes.RegisterRoutes(api, userUseCase)
		permissionsroutes.RegisterRoutes(api, permissionsUseCase)
		roleroutes.RegisterRoutes(api, roleUseCase)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
