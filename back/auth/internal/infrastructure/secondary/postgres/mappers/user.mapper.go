package mappers

import (
	"auth/internal/domain/user/userdtos"
	"auth/internal/infrastructure/secondary/postgres/models"
)

func MapToUserDTO(user *models.Users) *userdtos.UserDTO {
	return &userdtos.UserDTO{
		ID:       user.ID,
		Name:     user.Name,
		LastName: user.LastName,
	}
}

func MapToUserDTOs(users []models.Users) []userdtos.UserDTO {
	userDTOs := make([]userdtos.UserDTO, len(users))
	for i, user := range users {
		userDTOs[i] = *MapToUserDTO(&user)
	}
	return userDTOs
}

func MapToUserModel(userDTO *userdtos.UserDTO) *models.Users {
	return &models.Users{
		Name:     userDTO.Name,
		LastName: userDTO.LastName,
		Email:    userDTO.Email,
		Password: userDTO.Password,
	}
}

func MapToBusinessAndIdDTOs(businesses []models.Business) []userdtos.BusinessAndIdDTO {
	dtos := make([]userdtos.BusinessAndIdDTO, len(businesses))
	for i, b := range businesses {
		dtos[i] = userdtos.BusinessAndIdDTO{
			BusinessID:   b.ID,
			BusinessName: b.Name,
		}
	}
	return dtos
}
