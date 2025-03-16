package mappers

import (
	"auth/internal/domain/user/userdtos"
	"auth/internal/infrastructure/primary/handlers/handlers/userhandlers/dtos/response"
)

func MapLoginResponse(dto userdtos.LoginResponseDTO) response.LoginResponse {
	business := make([]response.BusinessAndIdResponse, len(dto.Business))
	for i, b := range dto.Business {
		business[i] = response.BusinessAndIdResponse{
			BusinessID:   b.BusinessID,
			BusinessName: b.BusinessName,
		}
	}

	permission := make([]response.PermissionResponse, len(dto.Permission))
	for i, p := range dto.Permission {
		permission[i] = response.PermissionResponse{
			PermissionID: p.PermissionID,
			UserID:       p.UserID,
			UserName:     p.UserName,
			RoleID:       p.RoleID,
			Role:         p.Role,
			Write:        p.Write,
			Read:         p.Read,
			Delete:       p.Delete,
			Update:       p.Update,
		}
	}

	return response.LoginResponse{
		AccessToken:  dto.Token,
		UserId:       dto.UserId,
		UserName:     dto.UserName,
		UserLastName: dto.UserLastName,
		RoleID:       dto.RoleID,
		RoleName:     dto.RoleName,
		Business:     business,
		Permission:   permission,
	}
}
