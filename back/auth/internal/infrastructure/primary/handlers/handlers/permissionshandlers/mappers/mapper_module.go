package mappers

import (
	"auth/internal/domain/permissions/permissionsdtos"
	"auth/internal/infrastructure/primary/handlers/handlers/permissionshandlers/dtos/response"
)

func MapToModuleResponse(dtos []permissionsdtos.ModuleDTO) []response.ModuleResponse {
	responses := make([]response.ModuleResponse, len(dtos))
	for i, dto := range dtos {
		responses[i] = response.ModuleResponse{
			ID:   dto.ID,
			Name: dto.Name,
		}
	}
	return responses
}
