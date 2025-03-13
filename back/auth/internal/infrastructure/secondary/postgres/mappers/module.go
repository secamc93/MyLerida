package mappers

import (
	"auth/internal/domain/permissions/permissionsdtos"
	"auth/internal/infrastructure/secondary/postgres/models"
)

// MapToModuleDTO transforma un modelo Module a ModuleDTO.
func MapToModuleDTO(mod models.Module) permissionsdtos.ModuleDTO {
	return permissionsdtos.ModuleDTO{
		ID:   mod.ID,
		Name: mod.Name,
	}
}

// MapToModuleDTOs transforma un slice de modelos Module a un slice de ModuleDTO.
func MapToModuleDTOs(mods []models.Module) []permissionsdtos.ModuleDTO {
	dtoList := make([]permissionsdtos.ModuleDTO, len(mods))
	for i, mod := range mods {
		dtoList[i] = MapToModuleDTO(mod)
	}
	return dtoList
}
