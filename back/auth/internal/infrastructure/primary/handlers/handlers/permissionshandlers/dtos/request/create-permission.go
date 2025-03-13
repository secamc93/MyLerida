package request

type CreatePermissionRequest struct {
	Name     string `json:"name" binding:"required"`
	ModuleID uint   `json:"module_id" binding:"required"`
}
