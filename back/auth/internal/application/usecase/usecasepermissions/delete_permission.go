package usecasepermissions

func (uc *PermissionsUseCase) DeletePermission(id uint) error {
	return uc.repo.DeletePermission(id)
}
