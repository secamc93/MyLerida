package mappers

import (
	"auth/internal/domain/user/dtos"
	"auth/internal/infrastructure/primary/handlers/handlers/userhandlers/dtos/response"
)

func MapLoginResponse(dto dtos.LoginResponseDTO) response.LoginResponse {
	return response.LoginResponse{
		AccessToken:  dto.Token,
		UserId:       dto.UserId,
		UserName:     dto.UserName,
		UserLastName: dto.UserLastName,
	}
}
