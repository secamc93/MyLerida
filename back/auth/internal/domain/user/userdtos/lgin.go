package userdtos

type LoginRequest struct {
	Email    string
	Password string
}

type LoginResponseDTO struct {
	Token        string
	UserId       uint
	UserName     string
	UserLastName string
	RoleID       uint
	RoleName     string
	Business     []BusinessAndIdDTO
	Permission   []PermissionDTO
}
type BusinessAndIdDTO struct {
	BusinessID   uint
	BusinessName string
}

type UserByEmailDTO struct {
	UserID       uint
	UserName     string
	UserLastName string
	RoleID       uint
	RoleName     string
}

type PermissionDTO struct {
	PermissionID uint
	UserID       uint
	UserName     string
	RoleID       uint
	Role         string
	Write        bool
	Read         bool
	Delete       bool
	Update       bool
}
