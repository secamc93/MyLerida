package mappers

import (
	"auth/internal/domain/user/dtos"
	"auth/internal/infrastructure/primary/handlers/dtos/request"
)

func MapToUserDTO(userRequest request.UserRequest) dtos.UserDTO {
	return dtos.UserDTO{
		ID:       userRequest.ID,
		Name:     userRequest.Name,
		Email:    userRequest.Email,
		Password: userRequest.Password,
	}
}
