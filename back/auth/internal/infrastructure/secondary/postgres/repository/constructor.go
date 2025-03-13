package repository

import (
	"auth/internal/domain/permissions/permissionsports"
	"auth/internal/domain/role/ports"
	userPorts "auth/internal/domain/user/ports"
	"auth/internal/infrastructure/secondary/postgres/connectpostgres"
	"auth/pkg/logger"
	"sync"
)

// Repository implementa múltiples interfaces de repositorio
type Repository struct {
	dbConnection connectpostgres.DBConnection
	log          logger.ILogger
}

var (
	instance *Repository
	once     sync.Once
)

// New crea una instancia única del repositorio
func New(db connectpostgres.DBConnection) *Repository {
	once.Do(func() {
		instance = &Repository{
			dbConnection: db,
			log:          logger.NewLogger(),
		}
	})
	return instance
}

func (r *Repository) NewUserRepository() userPorts.IUserRepository {
	return r
}

func (r *Repository) NewRoleRepository() ports.IRoleRepository {
	return r
}

func (r *Repository) NewPermissionRepository() permissionsports.IPermissionsRepository {
	return r
}
