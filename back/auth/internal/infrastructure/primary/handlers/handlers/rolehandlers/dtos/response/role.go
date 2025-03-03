package response

// Permission representa un permiso en la respuesta
type Permission struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Module Module `json:"module"`
}

// Module representa un módulo en la respuesta
type Module struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

// RoleResponse representa la respuesta de un rol
type RoleResponse struct {
	ID          uint         `json:"id"`
	Name        string       `json:"name"`
	Permissions []Permission `json:"permissions,omitempty"`
}
