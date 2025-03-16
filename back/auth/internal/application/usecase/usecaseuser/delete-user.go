package usecaseuser

import (
	"auth/internal/domain/user/usererrors"
)

func (u *UserUseCase) DeleteUser(id uint) error {
	user, err := u.repoUser.GetUserByID(id)
	if err != nil || user == nil {
		return usererrors.ErrUserNotFound
	}

	return u.repoUser.DeleteUser(id)
}
