package ports

import "auth/internal/domain/role/dtos"

type IRoleRepository interface {
	GetRoleByID(id uint) (*dtos.RoleDTO, error)
	ListRoles() ([]dtos.RoleDTO, error)
	CreateRole(roleDTO *dtos.RoleDTO) error
	DeleteRole(id uint) error
	UpdateRole(id uint, roleDTO *dtos.RoleDTO) error
	GetRoleByName(name string) (*dtos.RoleDTO, error)
}
