package usecaserole

import (
	"auth/internal/domain/role/dtos"
	roleErrors "auth/internal/domain/role/errors"
	"fmt"
)

// ListRoles devuelve todos los roles disponibles
func (u *RoleUseCase) ListRoles() ([]dtos.RoleDTO, error) {
	// Obtener los roles del repositorio
	roles, err := u.repo.ListRoles()
	if err != nil {
		u.log.Error("Error al listar los roles: %v", err)
		return nil, fmt.Errorf("%w: %v", roleErrors.ErrRoleNotFound, err)
	}

	// Verificar si hay roles
	if len(roles) == 0 {
		u.log.Info("No se encontraron roles")
		return []dtos.RoleDTO{}, nil
	}

	u.log.Info("Se encontraron %d roles", len(roles))
	return roles, nil
}
