package ports

import (
	"auth/internal/domain/user/dtos"
)

type IUserRepository interface {
	CreateUser(user *dtos.UserDTO) error
	GetUserByID(id uint) (*dtos.UserDTO, error)
	ListUsers() ([]dtos.UserDTO, error)
}
