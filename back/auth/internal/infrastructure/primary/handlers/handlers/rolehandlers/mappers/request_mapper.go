package mappers

import (
	"auth/internal/domain/role/dtos"
	"auth/internal/infrastructure/primary/handlers/handlers/rolehandlers/dtos/request"
)

// MapCreateRequestToRoleDTO convierte una solicitud de creación a un DTO del dominio
func MapCreateRequestToRoleDTO(req *request.CreateRoleRequest) *dtos.RoleDTO {
	if req == nil {
		return nil
	}

	return &dtos.RoleDTO{
		Name: req.Name,
		// Aquí se pueden mapear otros campos como Description si es necesario
		// Si se necesita mapear permisos, se podría hacer aquí también
	}
}

// MapUpdateRequestToRoleDTO convierte una solicitud de actualización a un DTO del dominio
func MapUpdateRequestToRoleDTO(req *request.UpdateRoleRequest, id uint) *dtos.RoleDTO {
	if req == nil {
		return nil
	}

	return &dtos.RoleDTO{
		ID:   id,
		Name: req.Name,
		// Aquí se pueden mapear otros campos como Description si es necesario
		// Si se necesita mapear permisos, se podría hacer aquí también
	}
}
