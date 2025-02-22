package usecaseuser

import (
	"auth/internal/domain/user/dtos"
	"auth/internal/domain/user/ports"
)

type IUserUseCase interface {
	CreateUser(userDTO *dtos.UserDTO) error
}

type UserUseCase struct {
	repo ports.IUserRepository
}

func New(repo ports.IUserRepository) IUserUseCase {
	return &UserUseCase{repo: repo}
}

func (u *UserUseCase) CreateUser(userDTO *dtos.UserDTO) error {
	return u.repo.CreateUser(userDTO)
}
