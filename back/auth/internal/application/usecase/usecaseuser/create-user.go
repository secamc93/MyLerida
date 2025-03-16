package usecaseuser

import (
	"auth/internal/domain/user/userdtos"
	"auth/internal/domain/user/usererrors"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

func (u *UserUseCase) CreateUser(userDTO *userdtos.UserDTO) error {

	if userDTO.Email == "" {
		return usererrors.ErrEmailEmpty
	}
	if userDTO.Name == "" {
		return usererrors.ErrNameEmpty
	}
	if userDTO.Password == "" {
		return usererrors.ErrPasswordEmpty
	}

	exists, err := u.repoUser.UserExistsByEmail(userDTO.Email)
	if err != nil {
		return err
	}
	if exists {
		return usererrors.ErrEmailAlreadyExists
	}

	if !isValidPassword(userDTO.Password) {
		return usererrors.ErrPasswordInvalid
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userDTO.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	userDTO.Password = string(hashedPassword)

	userDTO.Name = capitalizeWords(userDTO.Name)
	userDTO.LastName = capitalizeWords(userDTO.LastName)

	return u.repoUser.CreateUser(userDTO)
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
