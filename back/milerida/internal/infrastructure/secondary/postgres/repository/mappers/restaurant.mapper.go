package mappers

import (
	"my-lerida/internal/domain/restaurant/dtos"
	"my-lerida/internal/domain/restaurant/entities"
	"my-lerida/internal/infrastructure/secondary/postgres/models"

	"gorm.io/gorm"
)

func ToRestaurantDTO(restaurant models.Restaurant) dtos.RestaurantDTO {
	return dtos.RestaurantDTO{
		ID:           restaurant.ID,
		Name:         restaurant.Name,
		Address:      restaurant.Address,
		PhoneNumber:  restaurant.PhoneNumber,
		Email:        restaurant.Email,
		OpeningHours: restaurant.OpeningHours,
		CuisineType:  restaurant.CuisineType,
		AverageCost:  restaurant.AverageCost,
	}
}

func ToRestaurantModel(restaurantDTO dtos.RestaurantDTO) models.Restaurant {
	return models.Restaurant{
		Model:        gorm.Model{ID: restaurantDTO.ID},
		Name:         restaurantDTO.Name,
		Address:      restaurantDTO.Address,
		PhoneNumber:  restaurantDTO.PhoneNumber,
		Email:        restaurantDTO.Email,
		OpeningHours: restaurantDTO.OpeningHours,
		CuisineType:  restaurantDTO.CuisineType,
		AverageCost:  restaurantDTO.AverageCost,
	}
}

func ToRestaurantDTOList(restaurants []models.Restaurant) []dtos.RestaurantDTO {
	var restaurantDTOs []dtos.RestaurantDTO
	for _, restaurant := range restaurants {
		restaurantDTOs = append(restaurantDTOs, ToRestaurantDTO(restaurant))
	}
	return restaurantDTOs
}

func ToRestaurantModelList(restaurantDTOs []dtos.RestaurantDTO) []models.Restaurant {
	var restaurants []models.Restaurant
	for _, restaurantDTO := range restaurantDTOs {
		restaurants = append(restaurants, ToRestaurantModel(restaurantDTO))
	}
	return restaurants
}

func ToEntitiRestaurantModel(entitie *entities.RestaurantEntitie) *models.Restaurant {
	return &models.Restaurant{
		Name:         entitie.Name,
		Address:      entitie.Address,
		PhoneNumber:  entitie.PhoneNumber,
		Email:        entitie.Email,
		OpeningHours: entitie.OpeningHours,
		CuisineType:  entitie.CuisineType,
		AverageCost:  entitie.AverageCost,
	}
}

func ToEntitiRestaurantModelList(entities *[]entities.RestaurantEntitie) []models.Restaurant {
	var restaurantModels []models.Restaurant
	for _, entitie := range *entities {
		restaurantModels = append(restaurantModels, *ToEntitiRestaurantModel(&entitie))
	}
	return restaurantModels
}
