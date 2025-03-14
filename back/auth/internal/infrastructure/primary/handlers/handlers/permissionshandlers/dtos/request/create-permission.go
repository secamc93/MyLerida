package request

type CreatePermissionRequest struct {
	Write    bool `json:"write" binding:"required"`
	Read     bool `json:"read" binding:"required"`
	Update   bool `json:"update" binding:"required"`
	Delete   bool `json:"delete" binding:"required"`
	ModuleID uint `json:"module_id" binding:"required"`
}
