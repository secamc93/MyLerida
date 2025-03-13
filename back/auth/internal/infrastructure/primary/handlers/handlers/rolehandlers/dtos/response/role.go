package response

type Permission struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Module Module `json:"module"`
}

type Module struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type RoleResponse struct {
	ID          uint         `json:"id"`
	Name        string       `json:"name"`
	Permissions []Permission `json:"permissions,omitempty"`
}
