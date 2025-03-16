package repository

import (
	"auth/internal/domain/permissions/permissionsports"
	"auth/internal/domain/role/roleports"
	"auth/internal/domain/user/userports"
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

func (r *Repository) NewUserRepository() userports.IUserRepository {
	return r
}

func (r *Repository) NewRoleRepository() roleports.IRoleRepository {
	return r
}

func (r *Repository) NewPermissionRepository() permissionsports.IPermissionsRepository {
	return r
}
