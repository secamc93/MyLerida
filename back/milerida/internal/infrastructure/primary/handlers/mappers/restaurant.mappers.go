package mappers

import (
	"my-lerida/internal/domain/restaurant/dtos"
	"my-lerida/internal/domain/restaurant/entities"
	"my-lerida/internal/infrastructure/primary/handlers/request"
	"my-lerida/internal/infrastructure/primary/handlers/response"
)

func ToRestaurantDTO(r *dtos.RestaurantDTO) response.RestaurantResponse {
	return response.RestaurantResponse{
		ID:           r.ID,
		Name:         r.Name,
		Address:      r.Address,
		PhoneNumber:  r.PhoneNumber,
		Email:        r.Email,
		OpeningHours: r.OpeningHours,
		CuisineType:  r.CuisineType,
		AverageCost:  r.AverageCost,
	}
}

func ToRestaurantDTOList(restaurants []dtos.RestaurantDTO) []response.RestaurantResponse {
	var restaurantDTOs []response.RestaurantResponse
	for _, r := range restaurants {
		restaurantDTOs = append(restaurantDTOs, ToRestaurantDTO(&r))
	}
	return restaurantDTOs
}

func ToRestaurantEntitie(req *request.RestaurantRequest) entities.RestaurantEntitie {
	return entities.RestaurantEntitie{
		Name:         req.Name,
		Address:      req.Address,
		PhoneNumber:  req.PhoneNumber,
		Email:        req.Email,
		OpeningHours: req.OpeningHours,
		CuisineType:  req.CuisineType,
		AverageCost:  req.AverageCost,
	}
}

func ToRestaurantEntitieList(reqs []request.RestaurantRequest) []entities.RestaurantEntitie {
	var restaurantEntities []entities.RestaurantEntitie
	for _, req := range reqs {
		restaurantEntities = append(restaurantEntities, ToRestaurantEntitie(&req))
	}
	return restaurantEntities
}
