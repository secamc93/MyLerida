package mappers

import (
	"auth/internal/domain/role/dtos"
	"auth/internal/infrastructure/secondary/postgres/models"
)

// MapToRoleDTO convierte un modelo Role a un DTO
func MapToRoleDTO(role *models.Role) *dtos.RoleDTO {
	roleDTO := &dtos.RoleDTO{
		ID:   role.ID,
		Name: role.Name,
	}

	// Mapeamos los permisos si existen
	if len(role.Permissions) > 0 {
		permissions := make([]dtos.PermissionDTO, len(role.Permissions))
		for i, perm := range role.Permissions {
			permissions[i] = dtos.PermissionDTO{
				ID:     perm.ID,
				Write:  perm.Write,
				Read:   perm.Read,
				Update: perm.Update,
				Delete: perm.Delete,
				Module: dtos.ModuleDTO{
					ID:   perm.Module.ID,
					Name: perm.Module.Name,
				},
			}
		}
		roleDTO.Permissions = permissions
	}

	return roleDTO
}

// MapToRoleDTOs convierte una slice de modelos Role a una slice de DTOs
func MapToRoleDTOs(roles []models.Role) []dtos.RoleDTO {
	roleDTOs := make([]dtos.RoleDTO, len(roles))
	for i, role := range roles {
		roleDTOs[i] = *MapToRoleDTO(&role)
	}
	return roleDTOs
}

// MapToRoleModel convierte un DTO Role a un modelo
func MapToRoleModel(roleDTO *dtos.RoleDTO) *models.Role {
	if roleDTO == nil {
		return nil
	}

	role := &models.Role{
		Name: roleDTO.Name,
	}

	if roleDTO.ID > 0 {
		role.ID = roleDTO.ID
	}

	return role
}
