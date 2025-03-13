package usecaserole

import (
	"auth/internal/domain/role/roleerrors"
	"fmt"
)

func (u *RoleUseCase) DeleteRole(id uint) error {
	if id == 0 {
		return fmt.Errorf("el ID del rol no puede ser cero")
	}

	roleDTO, err := u.repo.GetRoleByID(id)
	if err != nil {
		u.log.Error("Error al verificar la existencia del rol con ID: %d, %v", id, err)
		return fmt.Errorf("%w: %v", roleerrors.ErrRoleNotFound, err)
	}

	if roleDTO == nil {
		return roleerrors.ErrRoleNotFound
	}

	err = u.repo.DeleteRole(id)
	if err != nil {
		u.log.Error("Error al eliminar el rol con ID: %d, %v", id, err)
		return fmt.Errorf("%w: %v", roleerrors.ErrRoleDeleteFailed, err)
	}

	u.log.Info("Rol eliminado con éxito: %s (ID: %d)", roleDTO.Name, id)
	return nil
}
