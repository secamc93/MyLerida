package usecaseuser

import (
	"auth/internal/domain/user/userdtos"
	"auth/internal/domain/user/usererrors"
)

func (u *UserUseCase) ListUsers() ([]userdtos.UserDTO, error) {
	users, err := u.repoUser.ListUsers()
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, usererrors.ErrNoUsersFound
	}
	return users, nil
}
