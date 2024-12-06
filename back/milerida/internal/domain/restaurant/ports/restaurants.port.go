package ports

import (
	"my-lerida/internal/domain/restaurant/dtos"
	"my-lerida/internal/domain/restaurant/entities"
)

type IRestaurantRepository interface {
	GetRestaurants(restaurant string) ([]dtos.RestaurantDTO, error)
	GetAllRestaurants() ([]dtos.RestaurantDTO, error)
	GetRestaurantByID(restaurantID uint) (dtos.RestaurantDTO, error)
	CreateRestaurant(restaurant entities.RestaurantEntitie) error
	UpdateRestaurant(restaurantID uint, restaurant entities.RestaurantEntitie) error
	DeleteRestaurant(restaurantID uint) error
}
