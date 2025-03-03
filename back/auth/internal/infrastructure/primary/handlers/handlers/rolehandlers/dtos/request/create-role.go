package request

// CreateRoleRequest representa la solicitud para crear un nuevo rol
// swagger:model CreateRoleRequest
type CreateRoleRequest struct {
	// Nombre del rol
	// required: true
	// example: Administrador
	Name string `json:"name" binding:"required"`

	// Descripción del rol
	// example: Rol con todos los permisos del sistema
	Description string `json:"description"`

	// IDs de los permisos asociados al rol
	// example: [1, 2, 3]
	Permissions []uint `json:"permissions"`
}
