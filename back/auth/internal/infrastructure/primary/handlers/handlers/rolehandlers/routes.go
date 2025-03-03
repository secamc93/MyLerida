package rolehandlers

import (
	"auth/internal/application/usecase/usecaserole"
	"auth/internal/infrastructure/primary/handlers/handlers/rolehandlers/handlers"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes registra todas las rutas relacionadas con los roles
func RegisterRoutes(r *gin.RouterGroup, roleUseCase usecaserole.IRoleUseCase) {
	handler := handlers.New(roleUseCase)

	roleRoutes := r.Group("/roles")
	{
		roleRoutes.GET("", handler.ListRoles)
		roleRoutes.GET("/:id", handler.GetRole)
		roleRoutes.POST("", handler.CreateRole)
		roleRoutes.PUT("/:id", handler.UpdateRole)
		roleRoutes.DELETE("/:id", handler.DeleteRole)
	}
}
