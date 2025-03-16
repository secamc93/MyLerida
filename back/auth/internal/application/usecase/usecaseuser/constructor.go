package usecaseuser

import (
	"auth/internal/domain/permissions/permissionsports"
	"auth/internal/domain/user/userdtos"
	"auth/internal/domain/user/userports"
	"auth/pkg/logger"
	"context"
)

type IUserUseCase interface {
	CreateUser(userDTO *userdtos.UserDTO) error
	Login(ctx context.Context, email, password string) (userdtos.LoginResponseDTO, error)
	ListUsers() ([]userdtos.UserDTO, error)
	DeleteUser(id uint) error
	UpdateUser(id uint, updatedDTO *userdtos.UserDTO) error
	GetUserByID(id uint) (*userdtos.UserDTO, error)
}

type UserUseCase struct {
	repoUser        userports.IUserRepository
	repoPermissions permissionsports.IPermissionsRepository
	log             logger.ILogger
}

func New(
	repoUser userports.IUserRepository,
	repoPermissions permissionsports.IPermissionsRepository,
) IUserUseCase {
	return &UserUseCase{
		repoUser:        repoUser,
		repoPermissions: repoPermissions,
		log:             logger.NewLogger(),
	}
}
