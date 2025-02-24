package errors

import "errors"

var (
	ErrEmailEmpty         = errors.New("el correo electrónico no puede estar vacío")
	ErrNameEmpty          = errors.New("el nombre no puede estar vacío")
	ErrPasswordEmpty      = errors.New("la contraseña no puede estar vacía")
	ErrPasswordInvalid    = errors.New("la contraseña debe contener al menos una letra mayúscula, una letra minúscula, un número y un carácter especial")
	ErrEmailAlreadyExists = errors.New("el correo electrónico ya existe")
	ErrUserNotFound       = errors.New("usuario no encontrado")
	ErrInvalidPassword    = errors.New("contraseña inválida")
	ErrJWTSecretNotSet    = errors.New("JWT secret no configurado")
)
