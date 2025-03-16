package userports

import (
	"auth/internal/domain/user/userdtos"
)

type IUserRepository interface {
	CreateUser(user *userdtos.UserDTO) error
	GetUserByID(id uint) (*userdtos.UserDTO, error)
	ListUsers() ([]userdtos.UserDTO, error)
	UserExistsByEmail(email string) (bool, error)
	GetPasswordByEmail(email string) (string, error)
	GetUserByEmail(email string) (*userdtos.UserByEmailDTO, error)
	DeleteUser(id uint) error
	UpdateUser(id uint, user *userdtos.UserDTO) error
	GetBusinessesByUserID(userID uint) ([]userdtos.BusinessAndIdDTO, error)
}
