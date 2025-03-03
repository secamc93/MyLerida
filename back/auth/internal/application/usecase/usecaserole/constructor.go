package usecaserole

import (
	"auth/internal/domain/role/dtos"
	"auth/internal/domain/role/ports"
	"auth/pkg/logger"
)

// IRoleUseCase define la interfaz para todos los casos de uso relacionados con roles
type IRoleUseCase interface {
	CreateRole(roleDTO *dtos.RoleDTO) (*dtos.RoleDTO, error)
	GetRoleByID(id uint) (*dtos.RoleDTO, error)
	ListRoles() ([]dtos.RoleDTO, error)
	DeleteRole(id uint) error
	UpdateRole(id uint, roleDTO *dtos.RoleDTO) (*dtos.RoleDTO, error)
}

type RoleUseCase struct {
	repo ports.IRoleRepository
	log  logger.ILogger
}

// New crea una nueva instancia de RoleUseCase
func New(repo ports.IRoleRepository) IRoleUseCase {
	return &RoleUseCase{
		repo: repo,
		log:  logger.NewLogger(),
	}
}
