package mappers

import (
	"auth/internal/domain/permissions/permissionsdtos"
	req "auth/internal/infrastructure/primary/handlers/handlers/permissionshandlers/dtos/request"
)

// MapCreatePermissionRequestToDTO transforma un CreatePermissionRequest a PermissionDTO.
func MapCreatePermissionRequestToDTO(r *req.CreatePermissionRequest) *permissionsdtos.PermissionDTO {
	return &permissionsdtos.PermissionDTO{
		ModuleID: r.ModuleID,
		Name:     r.Name,
		// ...mapear otros campos según corresponda...
	}
}
