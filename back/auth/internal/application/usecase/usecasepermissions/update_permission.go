package usecasepermissions

import (
	"auth/internal/domain/permissions/permissionsdtos"
)

func (uc *PermissionsUseCase) UpdatePermission(id uint, dto *permissionsdtos.PermissionDTO) (*permissionsdtos.PermissionDTO, error) {
	if err := uc.repo.UpdatePermission(id, dto); err != nil {
		return nil, err
	}
	return uc.repo.GetPermissionByID(id)
}
