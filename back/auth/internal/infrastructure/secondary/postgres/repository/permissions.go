package repository

import (
	"auth/internal/domain/permissions/permissionsdtos"
	"auth/internal/infrastructure/secondary/postgres/mappers"
	"auth/internal/infrastructure/secondary/postgres/models"

	"gorm.io/gorm"
)

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

func (r *Repository) GetPermissionsByUser(user uint) (*[]permissionsdtos.PermissionDTO, error) {
	var permissionDTOs []permissionsdtos.PermissionDTO
	err := r.dbConnection.GetDB().
		Table("users u").
		Select(`
			m.id,
			m."name",
			p.write,
			p.read,
			p.delete,
			p.update`).
		Joins("INNER JOIN roles r ON r.id = u.role_id").
		Joins("INNER JOIN role_permissions rp ON rp.role_id = r.id").
		Joins("INNER JOIN permissions p ON p.id = rp.permission_id").
		Joins("INNER JOIN modules m ON m.id = p.module_id").
		Where("u.id = ?", user).
		Scan(&permissionDTOs).Error

	if err != nil {
		return nil, err
	}

	return &permissionDTOs, nil
}
