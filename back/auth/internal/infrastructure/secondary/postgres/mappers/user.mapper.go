package mappers

import (
	"auth/internal/domain/user/dtos"
	"auth/internal/infrastructure/secondary/postgres/models"
)

func MapToUserDTO(user *models.Users) *dtos.UserDTO {
	return &dtos.UserDTO{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
}

func MapToUserDTOs(users []models.Users) []dtos.UserDTO {
	userDTOs := make([]dtos.UserDTO, len(users))
	for i, user := range users {
		userDTOs[i] = *MapToUserDTO(&user)
	}
	return userDTOs
}

func MapToUserModel(userDTO *dtos.UserDTO) *models.Users {
	return &models.Users{
		Name:     userDTO.Name,
		Email:    userDTO.Email,
		Password: userDTO.Password,
	}
}
