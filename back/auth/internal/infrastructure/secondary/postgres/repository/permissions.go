package repository

import (
	"auth/internal/domain/permissions/permissionsdtos"
	"auth/internal/infrastructure/secondary/postgres/mappers"
	"auth/internal/infrastructure/secondary/postgres/models"

	"gorm.io/gorm"
)

func (r *Repository) CreatePermission(dto *permissionsdtos.PermissionDTO) error {
	perm := mappers.MapToPermissionModel(dto)
	return r.dbConnection.GetDB().Create(perm).Error
}

func (r *Repository) GetPermissionByID(id uint) (*permissionsdtos.PermissionDTO, error) {
	var perm models.Permission
	if err := r.dbConnection.GetDB().
		Preload("Roles").
		Preload("Module").
		First(&perm, id).Error; err != nil {
		return nil, err
	}
	return mappers.MapToPermissionDTO(&perm), nil
}

func (r *Repository) UpdatePermission(id uint, dto *permissionsdtos.PermissionDTO) error {
	perm := mappers.MapToPermissionModel(dto)
	return r.dbConnection.GetDB().
		Model(&models.Permission{}).
		Where("id = ?", id).
		Updates(perm).Error
}

func (r *Repository) DeletePermission(id uint) error {
	perm := models.Permission{Model: gorm.Model{ID: id}}
	return r.dbConnection.GetDB().Delete(&perm).Error
}

func (r *Repository) GetModules() ([]permissionsdtos.ModuleDTO, error) {
	var modules []models.Module
	if err := r.dbConnection.GetDB().Find(&modules).Error; err != nil {
		return nil, err
	}
	return mappers.MapToModuleDTOs(modules), nil
}
