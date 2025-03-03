package usecaseuser

import (
	"auth/internal/domain/user/dtos"
	"auth/internal/domain/user/errors"
	"auth/pkg/env"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func (uc *UserUseCase) Login(email, password string) (dtos.LoginResponseDTO, error) {
	var response dtos.LoginResponseDTO
	user, err := uc.repo.GetUserByEmail(email)
	if err != nil {
		return response, errors.ErrUserNotFound
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return response, errors.ErrInvalidPassword
	}

	token, err := generateJWT(user.Email)
	if err != nil {
		return response, err
	}

	response = dtos.LoginResponseDTO{
		Token:        token,
		UserId:       user.ID,
		UserName:     user.Name,
		UserLastName: user.LastName,
	}

	return response, nil
}

func generateJWT(email string) (string, error) {
	secret := env.LoadEnv().JwtSecret
	if secret == "" {
		return "", errors.ErrJWTSecretNotSet
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
