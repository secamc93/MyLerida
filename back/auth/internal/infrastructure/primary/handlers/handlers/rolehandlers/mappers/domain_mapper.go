package mappers

import (
	"auth/internal/domain/role/dtos"
	"auth/internal/infrastructure/primary/handlers/handlers/rolehandlers/dtos/response"
)

// MapRoleDTOToResponse convierte un DTO del dominio a un DTO de respuesta
func MapRoleDTOToResponse(roleDTO *dtos.RoleDTO) *response.RoleResponse {
	if roleDTO == nil {
		return nil
	}

	roleResponse := &response.RoleResponse{
		ID:   roleDTO.ID,
		Name: roleDTO.Name,
	}

	// Mapeo de permisos si existen
	if len(roleDTO.Permissions) > 0 {
		permissions := make([]response.Permission, len(roleDTO.Permissions))
		for i, p := range roleDTO.Permissions {
			permissions[i] = response.Permission{
				ID:   p.ID,
				Name: p.Name,
				Module: response.Module{
					ID:   p.Module.ID,
					Name: p.Module.Name,
				},
			}
		}
		roleResponse.Permissions = permissions
	}

	return roleResponse
}

// MapRoleDTOsToResponseList convierte una lista de DTOs del dominio a una respuesta de lista
func MapRoleDTOsToResponseList(roleDTOs []dtos.RoleDTO) *response.RoleListResponse {
	response := &response.RoleListResponse{
		Roles: make([]response.RoleResponse, len(roleDTOs)),
	}

	for i, roleDTO := range roleDTOs {
		role := MapRoleDTOToResponse(&roleDTO)
		if role != nil {
			response.Roles[i] = *role
		}
	}

	return response
}
