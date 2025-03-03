package usecaserole

import (
	"auth/internal/domain/role/dtos"
	roleErrors "auth/internal/domain/role/errors"
	"fmt"
)

// GetRoleByID obtiene un rol por su ID
func (u *RoleUseCase) GetRoleByID(id uint) (*dtos.RoleDTO, error) {
	// Validar que el ID sea válido
	if id == 0 {
		return nil, fmt.Errorf("el ID del rol no puede ser cero")
	}

	// Obtener el rol del repositorio
	roleDTO, err := u.repo.GetRoleByID(id)
	if err != nil {
		u.log.Error("Error al obtener el rol con ID: %d, %v", id, err)
		return nil, fmt.Errorf("%w: %v", roleErrors.ErrRoleNotFound, err)
	}

	if roleDTO == nil {
		return nil, roleErrors.ErrRoleNotFound
	}

	u.log.Info("Rol obtenido con éxito: %s (ID: %d)", roleDTO.Name, roleDTO.ID)
	return roleDTO, nil
}
