package usecaserole

import (
	"auth/internal/domain/role/roledtos"
	"auth/internal/domain/role/roleerrors"
	"fmt"
)

func (u *RoleUseCase) ListRoles() ([]roledtos.RoleDTO, error) {
	roles, err := u.repo.ListRoles()
	if err != nil {
		u.log.Error("Error al listar los roles: %v", err)
		return nil, fmt.Errorf("%w: %v", roleerrors.ErrRoleNotFound, err)
	}

	if len(roles) == 0 {
		u.log.Info("No se encontraron roles")
		return []roledtos.RoleDTO{}, nil
	}

	u.log.Info("Se encontraron %d roles", len(roles))
	return roles, nil
}
