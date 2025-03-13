package usecaserole

import (
	"auth/internal/domain/role/dtos"
	"auth/internal/domain/role/roleerrors"
	"fmt"
)

func (u *RoleUseCase) CreateRole(roleDTO *dtos.RoleDTO) (*dtos.RoleDTO, error) {
	if roleDTO.Name == "" {
		return nil, roleerrors.ErrRoleNameEmpty
	}
	existingRole, err := u.repo.GetRoleByName(roleDTO.Name)
	if err == nil && existingRole != nil {
		return nil, roleerrors.ErrRoleAlreadyExists
	}
	err = u.repo.CreateRole(roleDTO)
	if err != nil {
		u.log.Error("Error al crear el rol")
		return nil, fmt.Errorf("%w: %v", roleerrors.ErrRoleCreationFailed, err)
	}
	createdRole, err := u.repo.GetRoleByName(roleDTO.Name)
	if err != nil {
		u.log.Error("Error al obtener el rol recién creado")
		return nil, fmt.Errorf("%w: %v", roleerrors.ErrRoleNotFound, err)
	}
	u.log.Info("Rol creado exitosamente: %s (ID: %d)", createdRole.Name, createdRole.ID)
	return createdRole, nil
}
