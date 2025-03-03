package ports

import (
	"auth/internal/domain/user/dtos"
)

type IUserRepository interface {
	CreateUser(user *dtos.UserDTO) error
	GetUserByID(id uint) (*dtos.UserDTO, error)
	ListUsers() ([]dtos.UserDTO, error)
	UserExistsByEmail(email string) (bool, error)
	GetPasswordByEmail(email string) (string, error)
	GetUserByEmail(email string) (*dtos.UserDTO, error)
	DeleteUser(id uint) error
	UpdateUser(id uint, user *dtos.UserDTO) error
}
