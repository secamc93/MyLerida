package repository

import (
	"auth/internal/domain/user/dtos"
	"auth/internal/domain/user/ports"
	"auth/internal/infrastructure/secondary/postgres"
	"auth/internal/infrastructure/secondary/postgres/mappers"
	"auth/internal/infrastructure/secondary/postgres/models"
	"sync"
)

type IUserRepository interface {
	CreateUser(user *models.Users) error
	GetUserByID(id uint) (*dtos.UserDTO, error)
	ListUsers() ([]dtos.UserDTO, error)
}

type UsersRepository struct {
	dbConnection postgres.DBConnection
}

var (
	instance *UsersRepository
	once     sync.Once
)

func New(db postgres.DBConnection) ports.IUserRepository {
	once.Do(func() {
		instance = &UsersRepository{
			dbConnection: db,
		}
	})
	return instance
}

func (r *UsersRepository) CreateUser(userDTO *dtos.UserDTO) error {
	user := mappers.MapToUserModel(userDTO)
	return r.dbConnection.GetDB().Create(user).Error
}

func (r *UsersRepository) GetUserByID(id uint) (*dtos.UserDTO, error) {
	var user models.Users
	if err := r.dbConnection.GetDB().First(&user, id).Error; err != nil {
		return nil, err
	}
	return mappers.MapToUserDTO(&user), nil
}

func (r *UsersRepository) ListUsers() ([]dtos.UserDTO, error) {
	var users []models.Users
	if err := r.dbConnection.GetDB().Find(&users).Error; err != nil {
		return nil, err
	}
	return mappers.MapToUserDTOs(users), nil
}
