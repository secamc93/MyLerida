package response

type PermissionResponse struct {
	BusinessesID   uint   `json:"businesses_id"`
	BusinessesName string `json:"businesses_name"`
	UserID         uint   `json:"user_id"`
	UserName       string `json:"user_name"`
	RoleID         uint   `json:"role_id"`
	Role           string `json:"role"`
	Write          bool   `json:"write"`
	Read           bool   `json:"read"`
	Delete         bool   `json:"delete"`
	Update         bool   `json:"update"`
}
