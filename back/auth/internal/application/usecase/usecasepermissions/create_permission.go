package usecasepermissions

import (
	"auth/internal/domain/permissions/permissionsdtos"
)

func (uc *PermissionsUseCase) CreatePermission(dto *permissionsdtos.PermissionDTO) (*permissionsdtos.PermissionDTO, error) {
	if err := uc.repo.CreatePermission(dto); err != nil {
		return nil, err
	}
	return dto, nil
}
