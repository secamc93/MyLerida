package usecaseuser

import (
	"auth/internal/domain/user/userdtos"
	"auth/internal/domain/user/usererrors"
	"auth/pkg/env"
	"context"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func (uc *UserUseCase) Login(ctx context.Context, email, password string) (userdtos.LoginResponseDTO, error) {
	var response userdtos.LoginResponseDTO
	user, err := uc.repoUser.GetUserByEmail(email)
	if err != nil {
		uc.log.Error("Error getting user by email : %s", email, err)
		return response, usererrors.ErrUserNotFound
	}
	userPassword, err := uc.repoUser.GetPasswordByEmail(email)
	if err != nil {
		uc.log.Error("Error getting password by email : %s", email, err)
		return response, usererrors.ErrUserNotFound
	}

	err = bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(password))
	if err != nil {
		uc.log.Error("Error comparing password %s", email, err)
		return response, usererrors.ErrInvalidPassword
	}

	businessUser, err := uc.repoUser.GetBusinessesByUserID(user.UserID)
	if err != nil {
		uc.log.Error("Error getting business by user id %d", user.UserID, err)
		return response, usererrors.ErrUserNotFound
	}

	permissions, err := uc.repoPermissions.GetPermissionsByUser(user.UserID)
	if err != nil {
		uc.log.Error("Error getting permissions by user id and business id", err)
		return response, usererrors.ErrUserNotFound
	}

	token, err := generateJWT(email)
	if err != nil {
		return response, err
	}

	var permissionsDtos []userdtos.PermissionDTO
	for _, permission := range *permissions {
		permissionsDtos = append(permissionsDtos, userdtos.PermissionDTO{
			PermissionID: permission.PermissionID,
			UserID:       permission.UserID,
			UserName:     permission.UserName,
			RoleID:       permission.RoleID,
			Role:         permission.Role,
			Write:        permission.Write,
			Read:         permission.Read,
			Delete:       permission.Delete,
			Update:       permission.Update,
		})
	}

	response = userdtos.LoginResponseDTO{
		Token:        token,
		UserId:       user.UserID,
		UserName:     user.UserName,
		UserLastName: user.UserLastName,
		RoleID:       user.RoleID,
		RoleName:     user.RoleName,
		Business:     businessUser,
		Permission:   permissionsDtos,
	}

	return response, nil
}

func generateJWT(email string) (string, error) {
	secret := env.LoadEnv().JwtSecret
	if secret == "" {
		return "", usererrors.ErrJWTSecretNotSet
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
