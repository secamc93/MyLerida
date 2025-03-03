package usecaserole

import (
	roleErrors "auth/internal/domain/role/errors"
	"fmt"
)

// DeleteRole elimina un rol por su ID
func (u *RoleUseCase) DeleteRole(id uint) error {
	// Validar que el ID sea válido
	if id == 0 {
		return fmt.Errorf("el ID del rol no puede ser cero")
	}

	// Verificar que el rol exista
	roleDTO, err := u.repo.GetRoleByID(id)
	if err != nil {
		u.log.Error("Error al verificar la existencia del rol con ID: %d, %v", id, err)
		return fmt.Errorf("%w: %v", roleErrors.ErrRoleNotFound, err)
	}

	if roleDTO == nil {
		return roleErrors.ErrRoleNotFound
	}

	// Eliminamos el rol
	err = u.repo.DeleteRole(id)
	if err != nil {
		u.log.Error("Error al eliminar el rol con ID: %d, %v", id, err)
		return fmt.Errorf("%w: %v", roleErrors.ErrRoleDeleteFailed, err)
	}

	u.log.Info("Rol eliminado con éxito: %s (ID: %d)", roleDTO.Name, id)
	return nil
}
