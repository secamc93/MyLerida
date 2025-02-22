package usecaseuser

import (
	"auth/internal/domain/user/dtos"
	"auth/internal/domain/user/errors"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

func (u *UserUseCase) CreateUser(userDTO *dtos.UserDTO) error {
	// Verificar si el email, nombre y contraseña no están vacíos
	if userDTO.Email == "" {
		return errors.ErrEmailEmpty
	}
	if userDTO.Name == "" {
		return errors.ErrNameEmpty
	}
	if userDTO.Password == "" {
		return errors.ErrPasswordEmpty
	}

	// Verificar si el email ya existe
	exists, err := u.repo.UserExistsByEmail(userDTO.Email)
	if err != nil {
		return err
	}
	if exists {
		return errors.ErrEmailAlreadyExists
	}

	// Validar la contraseña
	if !isValidPassword(userDTO.Password) {
		return errors.ErrPasswordInvalid
	}

	// Cifrar la contraseña
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userDTO.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	userDTO.Password = string(hashedPassword)

	return u.repo.CreateUser(userDTO)
}

func isValidPassword(password string) bool {
	var (
		hasMinLen      = len(password) >= 8
		hasUpper       = regexp.MustCompile(`[A-Z]`).MatchString(password)
		hasLower       = regexp.MustCompile(`[a-z]`).MatchString(password)
		hasNumber      = regexp.MustCompile(`[0-9]`).MatchString(password)
		hasSpecialChar = regexp.MustCompile(`[!@#~$%^&*()_+|<>?:{}]`).MatchString(password)
	)
	return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecialChar
}
