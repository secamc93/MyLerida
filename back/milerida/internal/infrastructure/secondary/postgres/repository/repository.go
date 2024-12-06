package repository

import (
	"my-lerida/internal/domain/restaurant/dtos"
	"my-lerida/internal/domain/restaurant/entities"
	"my-lerida/internal/domain/restaurant/ports"
	"my-lerida/internal/infrastructure/secondary/postgres"
	"my-lerida/internal/infrastructure/secondary/postgres/models"
	"my-lerida/internal/infrastructure/secondary/postgres/repository/mappers"
	"sync"
)

type RestaurantRepository interface {
	GetRestaurants(restaurant string) ([]dtos.RestaurantDTO, error)
	GetAllRestaurants() ([]dtos.RestaurantDTO, error)
	GetRestaurantByID(restaurantID uint) (dtos.RestaurantDTO, error)
	CreateRestaurant(restaurant entities.RestaurantEntitie) error
	UpdateRestaurant(restaurantID uint, restaurant entities.RestaurantEntitie) error
	DeleteRestaurant(restaurantID uint) error
}

type Restaurant struct {
	dbConnection postgres.DBConnection
}

var (
	instance *Restaurant
	once     sync.Once
)

func NewRestaurant(db postgres.DBConnection) ports.IRestaurantRepository {
	once.Do(func() {
		instance = &Restaurant{
			dbConnection: db,
		}
	})
	return instance
}

func (r *Restaurant) GetRestaurants(restaurant string) ([]dtos.RestaurantDTO, error) {
	var restaurantModels []models.Restaurant
	db := r.dbConnection.GetDB()
	err := db.Model(&restaurantModels).Where("name = ?", restaurant).Find(&restaurantModels).Error
	if err != nil {
		return nil, err
	}

	restaurantDTOS := mappers.ToRestaurantDTOList(restaurantModels)

	return restaurantDTOS, nil
}

func (r *Restaurant) GetAllRestaurants() ([]dtos.RestaurantDTO, error) {
	var restaurantModels []models.Restaurant
	db := r.dbConnection.GetDB()
	err := db.Model(&restaurantModels).Find(&restaurantModels).Error
	if err != nil {
		return nil, err
	}

	restaurantDTOS := mappers.ToRestaurantDTOList(restaurantModels)

	return restaurantDTOS, nil
}

func (r *Restaurant) GetRestaurantByID(restaurantID uint) (dtos.RestaurantDTO, error) {
	var restaurantModel models.Restaurant
	db := r.dbConnection.GetDB()

	err := db.Model(&restaurantModel).Find(restaurantModel).Error
	if err != nil {
		return dtos.RestaurantDTO{}, err
	}
	restaurantDTOS := mappers.ToRestaurantDTO(restaurantModel)

	return restaurantDTOS, nil

}

func (r *Restaurant) CreateRestaurant(restaurant entities.RestaurantEntitie) error {
	db := r.dbConnection.GetDB()
	err := db.Create(mappers.ToEntitiRestaurantModel(&restaurant)).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *Restaurant) UpdateRestaurant(restaurantID uint, restaurant entities.RestaurantEntitie) error {
	db := r.dbConnection.GetDB()
	restaurantModel := mappers.ToEntitiRestaurantModel(&restaurant)
	err := db.Model(&models.Restaurant{}).
		Where("id = ?", restaurantID).
		Updates(restaurantModel).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *Restaurant) DeleteRestaurant(restaurantID uint) error {
	db := r.dbConnection.GetDB()
	err := db.Delete(&models.Restaurant{}, restaurantID).Error
	if err != nil {
		return err
	}

	return nil
}
