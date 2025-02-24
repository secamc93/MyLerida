package usecaseuser

import (
	"auth/internal/domain/user/dtos"
	"auth/internal/domain/user/ports"
)

type IUserUseCase interface {
	CreateUser(userDTO *dtos.UserDTO) error
	Login(email, password string) (string, error)
}

type UserUseCase struct {
	repo ports.IUserRepository
}

func New(repo ports.IUserRepository) IUserUseCase {
	return &UserUseCase{repo: repo}
}
