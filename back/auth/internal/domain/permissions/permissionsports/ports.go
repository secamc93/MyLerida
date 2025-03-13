package permissionsports

import (
	"auth/internal/domain/permissions/permissionsdtos"
)

type IPermissionsRepository interface {
	CreatePermission(dto *permissionsdtos.PermissionDTO) error
	GetPermissionByID(id uint) (*permissionsdtos.PermissionDTO, error)
	UpdatePermission(id uint, dto *permissionsdtos.PermissionDTO) error
	DeletePermission(id uint) error
	GetModules() ([]permissionsdtos.ModuleDTO, error)
}
