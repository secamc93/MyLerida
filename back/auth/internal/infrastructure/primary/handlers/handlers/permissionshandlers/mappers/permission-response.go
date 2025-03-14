package mappers

import (
	"auth/internal/domain/permissions/permissionsdtos"
	"auth/internal/infrastructure/primary/handlers/handlers/permissionshandlers/dtos/response"
)

// MapToPermissionResponse mapea un PermissionDTO a un PermissionResponse
func MapToPermissionResponse(dto *permissionsdtos.PermissionDTO) *response.PermissionResponse {
	return &response.PermissionResponse{
		BusinessesID:   dto.BusinessesID,
		BusinessesName: dto.BusinessesName,
		UserID:         dto.UserID,
		UserName:       dto.UserName,
		RoleID:         dto.RoleID,
		Role:           dto.Role,
		Write:          dto.Write,
		Read:           dto.Read,
		Delete:         dto.Delete,
		Update:         dto.Update,
	}
}

// MapToPermissionResponseList mapea una lista de PermissionDTO a una lista de PermissionResponse
func MapToPermissionResponseList(dtos []permissionsdtos.PermissionDTO) []response.PermissionResponse {
	responses := make([]response.PermissionResponse, len(dtos))
	for i, dto := range dtos {
		responses[i] = *MapToPermissionResponse(&dto)
	}
	return responses
}
