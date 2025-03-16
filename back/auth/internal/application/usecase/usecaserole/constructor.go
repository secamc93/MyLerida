package usecaserole

import (
	"auth/internal/domain/role/roledtos"
	"auth/internal/domain/role/roleports"
	"auth/pkg/logger"
)

// IRoleUseCase define la interfaz para todos los casos de uso relacionados con roles
type IRoleUseCase interface {
	CreateRole(roleDTO *roledtos.RoleDTO) (*roledtos.RoleDTO, error)
	GetRoleByID(id uint) (*roledtos.RoleDTO, error)
	ListRoles() ([]roledtos.RoleDTO, error)
	DeleteRole(id uint) error
	UpdateRole(id uint, roleDTO *roledtos.RoleDTO) (*roledtos.RoleDTO, error)
}

type RoleUseCase struct {
	repo roleports.IRoleRepository
	log  logger.ILogger
}

// New crea una nueva instancia de RoleUseCase
func New(repo roleports.IRoleRepository) IRoleUseCase {
	return &RoleUseCase{
		repo: repo,
		log:  logger.NewLogger(),
	}
}
