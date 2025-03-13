package mappers

import (
	"auth/internal/domain/permissions/permissionsdtos"
	"auth/internal/infrastructure/secondary/postgres/models"
)

func MapToPermissionModel(dto *permissionsdtos.PermissionDTO) *models.Permission {
	return &models.Permission{
		ModuleID: dto.ModuleID,
		Name:     dto.Name,
	}
}

func MapToPermissionDTO(model *models.Permission) *permissionsdtos.PermissionDTO {
	return &permissionsdtos.PermissionDTO{
		Name: model.Name,
	}
}
