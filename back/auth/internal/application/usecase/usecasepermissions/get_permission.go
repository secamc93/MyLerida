package usecasepermissions

import (
	"auth/internal/domain/permissions/permissionsdtos"
)

func (uc *PermissionsUseCase) GetPermissionByID(id uint) (*permissionsdtos.PermissionDTO, error) {
	return uc.repo.GetPermissionByID(id)
}
