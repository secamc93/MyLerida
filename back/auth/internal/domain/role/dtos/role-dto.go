package dtos

type RoleDTO struct {
	ID          uint
	Name        string
	Permissions []PermissionDTO
}

type PermissionDTO struct {
	ID     uint
	Module ModuleDTO
	Write  bool
	Read   bool
	Update bool
	Delete bool
}

type ModuleDTO struct {
	ID   uint
	Name string
}
