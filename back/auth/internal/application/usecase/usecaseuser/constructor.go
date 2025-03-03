package usecaseuser

import (
	"auth/internal/domain/user/dtos"
	"auth/internal/domain/user/ports"
	"auth/pkg/logger"
)

type IUserUseCase interface {
	CreateUser(userDTO *dtos.UserDTO) error
	Login(email, password string) (dtos.LoginResponseDTO, error)
	ListUsers() ([]dtos.UserDTO, error)
	DeleteUser(id uint) error
	UpdateUser(id uint, updatedDTO *dtos.UserDTO) error
	GetUserByID(id uint) (*dtos.UserDTO, error)
}

type UserUseCase struct {
	repo ports.IUserRepository
	log  logger.ILogger
}

func New(repo ports.IUserRepository) IUserUseCase {
	return &UserUseCase{
		repo: repo,
		log:  logger.NewLogger(),
	}
}
