package restaurant

import (
	"my-lerida/internal/domain/restaurant/dtos"
	"my-lerida/internal/domain/restaurant/entities"
	"my-lerida/internal/domain/restaurant/ports"
)

type IRestaurantUseCase interface {
	GetRestaurants(restaurant string) ([]dtos.RestaurantDTO, error)
	CreateRestaurant(restaurant entities.RestaurantEntitie) error
	UpdateRestaurant(restaurantID uint, restaurant entities.RestaurantEntitie) error
	DeleteRestaurant(restaurantID uint) error
	GetRestaurantByID(RestaurantID uint) (dtos.RestaurantDTO, error)
}

type Restaurant struct {
	repo ports.IRestaurantRepository
}

func NewRestaurant(repo ports.IRestaurantRepository) IRestaurantUseCase {
	return &Restaurant{repo: repo}
}

func (r *Restaurant) GetRestaurants(restaurant string) ([]dtos.RestaurantDTO, error) {
	if restaurant == "" {
		return r.repo.GetAllRestaurants()
	}
	return r.repo.GetRestaurants(restaurant)

}

func (r *Restaurant) GetRestaurantByID(restaurantID uint) (dtos.RestaurantDTO, error) {
	return r.repo.GetRestaurantByID(restaurantID)
}

func (r *Restaurant) CreateRestaurant(restaurant entities.RestaurantEntitie) error {
	return r.repo.CreateRestaurant(restaurant)
}

func (r *Restaurant) UpdateRestaurant(restaurantID uint, restaurant entities.RestaurantEntitie) error {
	return r.repo.UpdateRestaurant(restaurantID, restaurant)
}

func (r *Restaurant) DeleteRestaurant(restaurantID uint) error {
	return r.repo.DeleteRestaurant(restaurantID)
}
