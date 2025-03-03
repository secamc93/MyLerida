package usecaserole

import (
	"auth/internal/domain/role/dtos"
	roleErrors "auth/internal/domain/role/errors"
	"fmt"
)

// UpdateRole actualiza un rol existente
func (u *RoleUseCase) UpdateRole(id uint, roleDTO *dtos.RoleDTO) (*dtos.RoleDTO, error) {
	// Validar que el ID sea válido
	if id == 0 {
		return nil, fmt.Errorf("el ID del rol no puede ser cero")
	}

	// Validar que el rol tenga un nombre
	if roleDTO.Name == "" {
		return nil, roleErrors.ErrRoleNameEmpty
	}

	// Verificar que el rol exista
	existingRole, err := u.repo.GetRoleByID(id)
	if err != nil {
		u.log.Error("Error al verificar la existencia del rol con ID: %d, %v", id, err)
		return nil, fmt.Errorf("%w: %v", roleErrors.ErrRoleNotFound, err)
	}

	if existingRole == nil {
		return nil, roleErrors.ErrRoleNotFound
	}

	// Verificar si existe otro rol con el mismo nombre pero diferente ID
	if roleDTO.Name != existingRole.Name {
		roleWithSameName, err := u.repo.GetRoleByName(roleDTO.Name)
		if err == nil && roleWithSameName != nil && roleWithSameName.ID != id {
			return nil, roleErrors.ErrRoleAlreadyExists
		}
	}

	// Asignamos el ID al DTO para la actualización
	roleDTO.ID = id

	// Actualizamos el rol
	err = u.repo.UpdateRole(id, roleDTO)
	if err != nil {
		u.log.Error("Error al actualizar el rol con ID: %d, %v", id, err)
		return nil, fmt.Errorf("%w: %v", roleErrors.ErrRoleUpdateFailed, err)
	}

	// Obtenemos el rol actualizado
	updatedRole, err := u.repo.GetRoleByID(id)
	if err != nil {
		u.log.Error("Error al obtener el rol actualizado con ID: %d, %v", id, err)
		return nil, fmt.Errorf("%w: %v", roleErrors.ErrRoleNotFound, err)
	}

	u.log.Info("Rol actualizado con éxito: %s (ID: %d)", updatedRole.Name, id)
	return updatedRole, nil
}
