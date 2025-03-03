package repository

import (
	"auth/internal/domain/role/dtos"
	"auth/internal/infrastructure/secondary/postgres/mappers"
	"auth/internal/infrastructure/secondary/postgres/models"
)

func (r *Repository) GetRoleByID(id uint) (*dtos.RoleDTO, error) {
	var role models.Role
	if err := r.dbConnection.GetDB().Preload("Permissions.Module").First(&role, id).Error; err != nil {
		return nil, err
	}
	roleDTO := mappers.MapToRoleDTO(&role)
	return roleDTO, nil
}

func (r *Repository) ListRoles() ([]dtos.RoleDTO, error) {
	var roles []models.Role
	if err := r.dbConnection.GetDB().Preload("Permissions.Module").Find(&roles).Error; err != nil {
		return nil, err
	}
	roleDTOs := mappers.MapToRoleDTOs(roles)
	return roleDTOs, nil
}

func (r *Repository) CreateRole(roleDTO *dtos.RoleDTO) error {
	role := mappers.MapToRoleModel(roleDTO)
	return r.dbConnection.GetDB().Create(role).Error
}

func (r *Repository) DeleteRole(id uint) error {
	role := models.Role{}
	role.ID = id
	return r.dbConnection.GetDB().Delete(&role).Error
}

func (r *Repository) UpdateRole(id uint, roleDTO *dtos.RoleDTO) error {
	role := mappers.MapToRoleModel(roleDTO)
	r.log.Info("Updating role with ID: %d", id)
	if err := r.dbConnection.GetDB().
		Model(&models.Role{}).
		Where("id = ?", id).
		Updates(role).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetRoleByName(name string) (*dtos.RoleDTO, error) {
	var role models.Role
	if err := r.dbConnection.GetDB().Preload("Permissions.Module").Where("name = ?", name).First(&role).Error; err != nil {
		return nil, err
	}
	roleDTO := mappers.MapToRoleDTO(&role)
	return roleDTO, nil
}
