package response

// RoleListResponse representa la respuesta para una lista de roles
type RoleListResponse struct {
	Roles []RoleResponse `json:"roles"`
}
