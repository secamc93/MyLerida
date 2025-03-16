package repository

import (
	"auth/internal/domain/user/userdtos"
	"auth/internal/infrastructure/secondary/postgres/mappers"
	"auth/internal/infrastructure/secondary/postgres/models"
)

func (r *Repository) GetUserByID(id uint) (*userdtos.UserDTO, error) {
	var user models.Users
	if err := r.dbConnection.GetDB().First(&user, id).Error; err != nil {
		return nil, err
	}
	return mappers.MapToUserDTO(&user), nil
}

func (r *Repository) ListUsers() ([]userdtos.UserDTO, error) {
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
func (r *Repository) GetUserByEmail(email string) (*userdtos.UserByEmailDTO, error) {
	var result userdtos.UserByEmailDTO
	err := r.dbConnection.GetDB().
		Table("users u").
		Select(`u.id as user_id,
				u."name" as user_name,
				u.last_name as user_last_name,
				r.id as role_id,
				r."name" as role_name`).
		Joins("INNER JOIN roles r ON r.id = u.role_id").
		Where("u.email = ?", email).
		Scan(&result).Error

	if err != nil {
		return nil, err
	}
	return &result, nil
}
func (r *Repository) CreateUser(userDTO *userdtos.UserDTO) error {
	user := mappers.MapToUserModel(userDTO)
	return r.dbConnection.GetDB().Create(user).Error
}

func (r *Repository) DeleteUser(id uint) error {
	user := models.Users{Model: models.Users{}.Model}
	user.ID = id
	return r.dbConnection.GetDB().Delete(&user).Error
}

func (r *Repository) UpdateUser(id uint, userDTO *userdtos.UserDTO) error {
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

func (r *Repository) GetBusinessesByUserID(userID uint) ([]userdtos.BusinessAndIdDTO, error) {
	var businesses []models.Business
	gdb := r.dbConnection.GetDB().Table("businesses").
		Select("businesses.id, businesses.name").
		Joins("INNER JOIN user_businesses ON user_businesses.business_id = businesses.id").
		Joins("INNER JOIN users ON users.id = user_businesses.users_id").
		Where("users.id = ?", userID).
		Find(&businesses)
	if gdb.Error != nil {
		return nil, gdb.Error
	}
	return mappers.MapToBusinessAndIdDTOs(businesses), nil
}
