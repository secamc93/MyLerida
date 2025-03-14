package usecasepermissions

import (
	"auth/internal/domain/permissions/permissionsdtos"
	"auth/internal/domain/permissions/permissionsports"
	"auth/pkg/logger"
	"context"
)

type IPermissionsUseCase interface {
	GetModules() ([]permissionsdtos.ModuleDTO, error)
	GetPermissionByBussinesAndUser(ctx context.Context, bussinesID uint, userID uint) (*[]permissionsdtos.PermissionDTO, error)
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
