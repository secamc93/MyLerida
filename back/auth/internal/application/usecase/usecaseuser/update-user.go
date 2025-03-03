package usecaseuser

import (
	"auth/internal/domain/user/dtos"
	"auth/internal/domain/user/errors"
	"strings"

	"github.com/rs/zerolog/log"
)

func (u *UserUseCase) UpdateUser(id uint, updatedDTO *dtos.UserDTO) error {

	updatedDTO.Name = capitalizeWords(updatedDTO.Name)
	updatedDTO.LastName = capitalizeWords(updatedDTO.LastName)

	existing, err := u.repo.GetUserByID(id)
	if err != nil || existing == nil {
		log.Error().Err(err).Msg("Error getting user by ID")
		return errors.ErrUserNotFound
	}

	if updatedDTO.Email != existing.Email {
		exists, err := u.repo.UserExistsByEmail(updatedDTO.Email)
		if err != nil {
			log.Error().Err(err).Msg("Error checking if email exists")
			return err
		}
		if exists {
			log.Error().Msg("Email already exists")
			return errors.ErrEmailAlreadyExists
		}
	}

	err = u.repo.UpdateUser(id, updatedDTO)
	if err != nil {
		log.Error().Err(err).Msg("Error updating user")
		return err
	}

	return nil
}

func capitalizeWords(s string) string {
	return strings.Title(strings.ToLower(s))
}
