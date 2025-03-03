package request

// UpdateRoleRequest representa la solicitud para actualizar un rol existente
type UpdateRoleRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Permissions []uint `json:"permissions"`
}
