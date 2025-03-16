package roleports

import "auth/internal/domain/role/roledtos"

type IRoleRepository interface {
	GetRoleByID(id uint) (*roledtos.RoleDTO, error)
	ListRoles() ([]roledtos.RoleDTO, error)
	CreateRole(roleDTO *roledtos.RoleDTO) error
	DeleteRole(id uint) error
	UpdateRole(id uint, roleDTO *roledtos.RoleDTO) error
	GetRoleByName(name string) (*roledtos.RoleDTO, error)
}
