package repository

import (
	"auth/internal/domain/user/dtos"
	"auth/internal/infrastructure/secondary/postgres/mappers"
	"auth/internal/infrastructure/secondary/postgres/models"
)

func (r *Repository) GetUserByID(id uint) (*dtos.UserDTO, error) {
	var user models.Users
	if err := r.dbConnection.GetDB().First(&user, id).Error; err != nil {
		return nil, err
	}
	return mappers.MapToUserDTO(&user), nil
}

func (r *Repository) ListUsers() ([]dtos.UserDTO, error) {
	var users []models.Users
	if err := r.dbConnection.GetDB().Find(&users).Error; err != nil {
		return nil, err
	}
	return mappers.MapToUserDTOs(users), nil
}
func (r *Repository) UserExistsByEmail(email string) (bool, error) {
	var user models.Users
	if err := r.dbConnection.GetDB().Where("email = ?", email).First(&user).Error; err != nil {
		if err.Error() == "record not found" {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
func (r *Repository) GetPasswordByEmail(email string) (string, error) {
	var user models.Users
	if err := r.dbConnection.GetDB().Debug().Where("email = ?", email).First(&user).Error; err != nil {

		return "", err
	}
	return user.Password, nil
}
func (r *Repository) GetUserByEmail(email string) (*dtos.UserDTO, error) {
	var user models.Users
	if err := r.dbConnection.GetDB().Debug().Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return mappers.MapToUserDTO(&user), nil
}
func (r *Repository) CreateUser(userDTO *dtos.UserDTO) error {
	user := mappers.MapToUserModel(userDTO)
	return r.dbConnection.GetDB().Create(user).Error
}

func (r *Repository) DeleteUser(id uint) error {
	user := models.Users{Model: models.Users{}.Model}
	user.ID = id
	return r.dbConnection.GetDB().Delete(&user).Error
}

func (r *Repository) UpdateUser(id uint, userDTO *dtos.UserDTO) error {
	user := mappers.MapToUserModel(userDTO)
	r.log.Info("Updating user with ID: %s ", user.LastName)
	if err := r.dbConnection.GetDB().Debug().
		Model(&models.Users{}).
		Where("id = ?", id).
		Updates(user).Error; err != nil {
		return err
	}

	return nil
}
