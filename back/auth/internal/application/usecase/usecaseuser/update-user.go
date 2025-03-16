package usecaseuser

import (
	"auth/internal/domain/user/userdtos"
	"auth/internal/domain/user/usererrors"
	"strings"

	"github.com/rs/zerolog/log"
)

func (u *UserUseCase) UpdateUser(id uint, updatedDTO *userdtos.UserDTO) error {

	updatedDTO.Name = capitalizeWords(updatedDTO.Name)
	updatedDTO.LastName = capitalizeWords(updatedDTO.LastName)

	existing, err := u.repoUser.GetUserByID(id)
	if err != nil || existing == nil {
		log.Error().Err(err).Msg("Error getting user by ID")
		return usererrors.ErrUserNotFound
	}

	if updatedDTO.Email != existing.Email {
		exists, err := u.repoUser.UserExistsByEmail(updatedDTO.Email)
		if err != nil {
			log.Error().Err(err).Msg("Error checking if email exists")
			return err
		}
		if exists {
			log.Error().Msg("Email already exists")
			return usererrors.ErrEmailAlreadyExists
		}
	}

	err = u.repoUser.UpdateUser(id, updatedDTO)
	if err != nil {
		log.Error().Err(err).Msg("Error updating user")
		return err
	}

	return nil
}

func capitalizeWords(s string) string {
	return strings.Title(strings.ToLower(s))
}
