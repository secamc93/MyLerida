package permissionsports

import (
	"auth/internal/domain/permissions/permissionsdtos"
)

type IPermissionsRepository interface {
	GetBusinessPermissions(business uint) (*[]permissionsdtos.PermissionDTO, error)
	GetBusinessPermissionsByUser(business uint, user uint) (*[]permissionsdtos.PermissionDTO, error)
	GetModules() ([]permissionsdtos.ModuleDTO, error)
}
