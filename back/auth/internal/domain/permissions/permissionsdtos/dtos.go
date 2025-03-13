package permissionsdtos

type PermissionDTO struct {
	ID       uint
	ModuleID uint
	Name     string
}

type ModuleDTO struct {
	ID   uint
	Name string
}
