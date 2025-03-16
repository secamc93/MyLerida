package permissionsports

import (
	"auth/internal/domain/permissions/permissionsdtos"
)

type IPermissionsRepository interface {
	GetPermissionsByUser(user uint) (*[]permissionsdtos.PermissionDTO, error)
	GetModules() ([]permissionsdtos.ModuleDTO, error)
}
