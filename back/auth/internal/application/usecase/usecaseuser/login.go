package usecaseuser

import (
	"auth/internal/domain/user/errors"
	"auth/pkg/env"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func (uc *UserUseCase) Login(email, password string) (string, error) {
	user, err := uc.repo.GetUserByEmail(email)
	if err != nil {
		return "", errors.ErrUserNotFound
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.ErrInvalidPassword
	}

	token, err := generateJWT(user.Email)
	if err != nil {
		return "", err
	}

	return token, nil
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
