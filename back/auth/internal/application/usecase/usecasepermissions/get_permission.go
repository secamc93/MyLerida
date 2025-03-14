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

	if userID == 0 {
		permissions, err := uc.repo.GetBusinessPermissions(bussinesID)
		if err != nil {
			uc.log.Error(err.Error())
			return nil, err
		}
		return permissions, nil
	} else {
		permissions, err := uc.repo.GetBusinessPermissionsByUser(bussinesID, userID)
		if err != nil {
			uc.log.Error(err.Error())
			return nil, err
		}
		return permissions, nil
	}
}
