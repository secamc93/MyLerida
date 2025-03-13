package usecaserole

import (
	"auth/internal/domain/role/dtos"
	"auth/internal/domain/role/roleerrors"
	"fmt"
)

func (u *RoleUseCase) GetRoleByID(id uint) (*dtos.RoleDTO, error) {
	if id == 0 {
		return nil, fmt.Errorf("el ID del rol no puede ser cero")
	}

	roleDTO, err := u.repo.GetRoleByID(id)
	if err != nil {
		u.log.Error("Error al obtener el rol con ID: %d, %v", id, err)
		return nil, fmt.Errorf("%w: %v", roleerrors.ErrRoleNotFound, err)
	}

	if roleDTO == nil {
		return nil, roleerrors.ErrRoleNotFound
	}

	u.log.Info("Rol obtenido con éxito: %s (ID: %d)", roleDTO.Name, roleDTO.ID)
	return roleDTO, nil
}
