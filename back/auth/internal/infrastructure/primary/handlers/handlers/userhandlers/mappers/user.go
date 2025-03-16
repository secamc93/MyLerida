package mappers

import (
	"auth/internal/domain/user/userdtos"
	"auth/internal/infrastructure/primary/handlers/handlers/userhandlers/dtos/request"
	"auth/internal/infrastructure/primary/handlers/handlers/userhandlers/dtos/response"
)

func MapToUserDTO(userRequest request.UserRequest) userdtos.UserDTO {
	return userdtos.UserDTO{
		Name:     userRequest.Name,
		Email:    userRequest.Email,
		Password: userRequest.Password,
	}
}
func MapUserResponse(dto userdtos.UserDTO) response.UserResponse {
	return response.UserResponse{
		ID:       dto.ID,
		Name:     dto.Name,
		LastName: dto.LastName,
	}
}

func MapUserResponses(dtos []userdtos.UserDTO) []response.UserResponse {
	responses := make([]response.UserResponse, len(dtos))
	for i, dto := range dtos {
		responses[i] = MapUserResponse(dto)
	}
	return responses
}

func ToUserUpdateDTO(req request.UserUpdateRequest) userdtos.UserDTO {
	return userdtos.UserDTO{
		Name:     req.Name,
		LastName: req.LastName,
		Email:    req.Email,
	}
}
