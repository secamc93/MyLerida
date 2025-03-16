package usecaseuser

import (
	"auth/internal/domain/user/userdtos"
	"auth/internal/domain/user/usererrors"
)

func (u *UserUseCase) GetUserByID(id uint) (*userdtos.UserDTO, error) {
	user, err := u.repoUser.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, usererrors.ErrUserNotFound
	}
	return user, nil
}
