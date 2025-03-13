package userroutes

import (
	"auth/internal/application/usecase/usecaseuser"
	"auth/internal/infrastructure/primary/handlers/handlers/userhandlers/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.RouterGroup, userUseCase usecaseuser.IUserUseCase) {
	handler := handlers.New(userUseCase)
	userRoutes := r.Group("users")
	{
		userRoutes.POST("/", handler.CreateUser)
		userRoutes.GET("/", handler.GetUsers)
		userRoutes.GET("/:id", handler.GetUser)
		userRoutes.PUT("/:id", handler.UpdateUser)
		userRoutes.DELETE("/:id", handler.DeleteUser)
		userRoutes.POST("/login", handler.Login)
	}
}
