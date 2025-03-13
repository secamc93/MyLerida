package usecasepermissions

import (
	"auth/internal/domain/permissions/permissionsdtos"
)

func (uc *PermissionsUseCase) GetModules() ([]permissionsdtos.ModuleDTO, error) {
	return uc.repo.GetModules()
}
