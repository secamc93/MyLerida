package permissionsroutes

import (
	"auth/internal/application/usecase/usecasepermissions"
	"auth/internal/infrastructure/primary/handlers/handlers/permissionshandlers/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.RouterGroup, peromossionsUseCase usecasepermissions.IPermissionsUseCase) {
	handler := handlers.New(peromossionsUseCase)
	permissionsRoutes := r.Group("/permissions")
	{
		permissionsRoutes.GET("/modules", handler.GetModules)
		permissionsRoutes.POST("/create", handler.CreatePermission)
	}
}
