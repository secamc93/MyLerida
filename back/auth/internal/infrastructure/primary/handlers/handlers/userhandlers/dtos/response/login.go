package response

type LoginResponse struct {
	AccessToken  string                  `json:"accessToken"`
	UserId       uint                    `json:"userId"`
	UserName     string                  `json:"userName"`
	UserLastName string                  `json:"userLastName"`
	RoleID       uint                    `json:"roleId"`
	RoleName     string                  `json:"roleName"`
	Business     []BusinessAndIdResponse `json:"business"`
	Permission   []PermissionResponse    `json:"permission"`
}

type BusinessAndIdResponse struct {
	BusinessID   uint   `json:"businessId"`
	BusinessName string `json:"businessName"`
}

type UserByEmailResponse struct {
	UserID       uint   `json:"userId"`
	UserName     string `json:"userName"`
	UserLastName string `json:"userLastName"`
	RoleID       uint   `json:"roleId"`
	RoleName     string `json:"roleName"`
}

type PermissionResponse struct {
	PermissionID uint   `json:"permissionId"`
	UserID       uint   `json:"userId"`
	UserName     string `json:"userName"`
	RoleID       uint   `json:"roleId"`
	Role         string `json:"role"`
	Write        bool   `json:"write"`
	Read         bool   `json:"read"`
	Delete       bool   `json:"delete"`
	Update       bool   `json:"update"`
}
