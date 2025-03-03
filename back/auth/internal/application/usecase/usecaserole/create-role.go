package usecaserole

import (
	"auth/internal/domain/role/dtos"
	roleErrors "auth/internal/domain/role/errors"
	"fmt"
)

func (u *RoleUseCase) CreateRole(roleDTO *dtos.RoleDTO) (*dtos.RoleDTO, error) {
	// Validar que el rol tenga un nombre
	if roleDTO.Name == "" {
		return nil, roleErrors.ErrRoleNameEmpty
	}

	// Validar que el rol no exista previamente
	existingRole, err := u.repo.GetRoleByName(roleDTO.Name)
	if err == nil && existingRole != nil {
		return nil, roleErrors.ErrRoleAlreadyExists
	}

	// Crear el nuevo rol
	err = u.repo.CreateRole(roleDTO)
	if err != nil {
		u.log.Error("Error al crear el rol")
		return nil, fmt.Errorf("%w: %v", roleErrors.ErrRoleCreationFailed, err)
	}

	// Obtener el rol creado para devolverlo con su ID asignado
	createdRole, err := u.repo.GetRoleByName(roleDTO.Name)
	if err != nil {
		u.log.Error("Error al obtener el rol recién creado")
		return nil, fmt.Errorf("%w: %v", roleErrors.ErrRoleNotFound, err)
	}

	u.log.Info("Rol creado exitosamente: %s (ID: %d)", createdRole.Name, createdRole.ID)
	return createdRole, nil
}
