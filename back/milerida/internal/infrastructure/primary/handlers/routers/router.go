package routers

import (
	"my-lerida/internal/application/usecase/restaurant"
	restaurantHandler "my-lerida/internal/infrastructure/primary/handlers"
	"my-lerida/internal/infrastructure/secondary/postgres"
	"my-lerida/internal/infrastructure/secondary/postgres/repository"
	"os"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	dbConnection := postgres.NewDBConnection()
	restaurantRepo := repository.NewRestaurant(dbConnection)
	restaurantUseCase := restaurant.NewRestaurant(restaurantRepo)

	urlBase := os.Getenv("URL_BASE")
	if urlBase == "" {
		urlBase = "/api"
	}

	api := router.Group(urlBase)
	{
		handler := restaurantHandler.NewRestaurantHandler(restaurantUseCase)
		api.GET("/restaurants", handler.GetRestaurants)
		api.POST("/restaurants", handler.CreateRestaurantHandler)
		api.GET("/restaurants/:id", handler.GetRestaurantByID)
		api.PUT("/restaurants/:id", handler.UpdateRestaurantHandler)
		api.DELETE("/restaurants/:id", handler.DeleteRestaurant)
	}

	return router
}
