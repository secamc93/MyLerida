package usecaseuser

import (
	"auth/internal/domain/user/errors"
)

func (u *UserUseCase) DeleteUser(id uint) error {
	user, err := u.repo.GetUserByID(id)
	if err != nil || user == nil {
		return errors.ErrUserNotFound
	}

	return u.repo.DeleteUser(id)
}
