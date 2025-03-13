package usecasepermissions

import (
	"auth/internal/domain/permissions/permissionsdtos"
	"auth/internal/domain/permissions/permissionsports"
	"auth/pkg/logger"
)

type IPermissionsUseCase interface {
	CreatePermission(dto *permissionsdtos.PermissionDTO) (*permissionsdtos.PermissionDTO, error)
	GetPermissionByID(id uint) (*permissionsdtos.PermissionDTO, error)
	UpdatePermission(id uint, dto *permissionsdtos.PermissionDTO) (*permissionsdtos.PermissionDTO, error)
	DeletePermission(id uint) error
	GetModules() ([]permissionsdtos.ModuleDTO, error)
}

type PermissionsUseCase struct {
	repo permissionsports.IPermissionsRepository
	log  logger.ILogger
}

func New(repo permissionsports.IPermissionsRepository) IPermissionsUseCase {
	return &PermissionsUseCase{
		repo: repo,
		log:  logger.NewLogger(),
	}
}
