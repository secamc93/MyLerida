package dtos

type RoleDTO struct {
	ID          uint
	Name        string
	Permissions []PermissionDTO
}

type PermissionDTO struct {
	ID     uint
	Name   string
	Module ModuleDTO
}

type ModuleDTO struct {
	ID   uint
	Name string
}
