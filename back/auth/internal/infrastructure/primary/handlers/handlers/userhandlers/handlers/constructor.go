package handlers

import (
	"auth/internal/application/usecase/usecaseuser"
	"auth/pkg/logger"
)

type UserHandler struct {
	useCase usecaseuser.IUserUseCase
	log     logger.ILogger
}

func New(useCase usecaseuser.IUserUseCase) *UserHandler {
	return &UserHandler{
		useCase: useCase,
		log:     logger.NewLogger(),
	}
}
