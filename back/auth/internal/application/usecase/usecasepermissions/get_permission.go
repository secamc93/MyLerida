package usecasepermissions

import (
	"auth/internal/domain/permissions/permissionsdtos"
	"context"
)

func (uc *PermissionsUseCase) GetPermissionByBussinesAndUser(
	ctx context.Context,
	bussinesID uint,
	userID uint,
) (*[]permissionsdtos.PermissionDTO, error) {
	return nil, nil
}
