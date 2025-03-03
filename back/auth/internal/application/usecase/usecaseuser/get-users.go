package usecaseuser

import (
	"auth/internal/domain/user/dtos"
	"auth/internal/domain/user/errors"
)

func (u *UserUseCase) ListUsers() ([]dtos.UserDTO, error) {
	users, err := u.repo.ListUsers()
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, errors.ErrNoUsersFound
	}
	return users, nil
}
