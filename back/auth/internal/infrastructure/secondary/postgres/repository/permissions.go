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
func (r *Repository) GetBusinessPermissions(business uint) (*[]permissionsdtos.PermissionDTO, error) {
	var permissionDTOs []permissionsdtos.PermissionDTO
	err := r.dbConnection.GetDB().
		Table("users u").
		Select(`
			b.id AS businesses_id,
			b.name AS businesses_name,
			u.id AS user_id,
			u.name AS user_name,
			r.id AS role_id,
			r.name AS role,
			p.write,
			p.read,
			p.delete,
			p.update`).
		Joins("INNER JOIN roles r ON r.id = u.role_id").
		Joins("INNER JOIN role_permissions rp ON rp.role_id = r.id").
		Joins("INNER JOIN permissions p ON p.id = rp.permission_id").
		Joins("INNER JOIN user_businesses ub ON ub.users_id = u.id").
		Joins("INNER JOIN businesses b ON b.id = ub.business_id").
		Where("b.id = ?", business).
		Scan(&permissionDTOs).Error
	if err != nil {
		return nil, err
	}

	return &permissionDTOs, nil
}
func (r *Repository) GetBusinessPermissionsByUser(business uint, user uint) (*[]permissionsdtos.PermissionDTO, error) {
	var permissionDTOs []permissionsdtos.PermissionDTO
	err := r.dbConnection.GetDB().
		Table("users u").
		Select(`
			b.id AS businesses_id,
			b.name AS businesses_name,
			u.id AS user_id,
			u.name AS user_name,
			r.id AS role_id,
			r.name AS role,
			p.write,
			p.read,
			p.delete,
			p.update`).
		Joins("INNER JOIN roles r ON r.id = u.role_id").
		Joins("INNER JOIN role_permissions rp ON rp.role_id = r.id").
		Joins("INNER JOIN permissions p ON p.id = rp.permission_id").
		Joins("INNER JOIN user_businesses ub ON ub.users_id = u.id").
		Joins("INNER JOIN businesses b ON b.id = ub.business_id").
		Where("b.id = ? AND u.id = ?", business, user).
		Scan(&permissionDTOs).Error

	if err != nil {
		return nil, err
	}

	return &permissionDTOs, nil
}
