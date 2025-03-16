package usecaserole

import (
	"auth/internal/domain/role/roledtos"
	"auth/internal/domain/role/roleerrors"
	"fmt"
)

func (u *RoleUseCase) UpdateRole(id uint, roleDTO *roledtos.RoleDTO) (*roledtos.RoleDTO, error) {
	if id == 0 {
		return nil, fmt.Errorf("el ID del rol no puede ser cero")
	}

	if roleDTO.Name == "" {
		return nil, roleerrors.ErrRoleNameEmpty
	}

	existingRole, err := u.repo.GetRoleByID(id)
	if err != nil {
		u.log.Error("Error al verificar la existencia del rol con ID: %d, %v", id, err)
		return nil, fmt.Errorf("%w: %v", roleerrors.ErrRoleNotFound, err)
	}

	if existingRole == nil {
		return nil, roleerrors.ErrRoleNotFound
	}

	if roleDTO.Name != existingRole.Name {
		roleWithSameName, err := u.repo.GetRoleByName(roleDTO.Name)
		if err == nil && roleWithSameName != nil && roleWithSameName.ID != id {
			return nil, roleerrors.ErrRoleAlreadyExists
		}
	}

	roleDTO.ID = id

	err = u.repo.UpdateRole(id, roleDTO)
	if err != nil {
		u.log.Error("Error al actualizar el rol con ID: %d, %v", id, err)
		return nil, fmt.Errorf("%w: %v", roleerrors.ErrRoleUpdateFailed, err)
	}

	updatedRole, err := u.repo.GetRoleByID(id)
	if err != nil {
		u.log.Error("Error al obtener el rol actualizado con ID: %d, %v", id, err)
		return nil, fmt.Errorf("%w: %v", roleerrors.ErrRoleNotFound, err)
	}

	u.log.Info("Rol actualizado con éxito: %s (ID: %d)", updatedRole.Name, id)
	return updatedRole, nil
}
