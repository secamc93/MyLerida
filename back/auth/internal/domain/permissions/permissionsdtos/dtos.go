package permissionsdtos

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

type ModuleDTO struct {
	ID   uint
	Name string
}
