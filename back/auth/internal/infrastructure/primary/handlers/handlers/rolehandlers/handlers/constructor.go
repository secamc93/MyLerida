package handlers

import (
	"auth/internal/application/usecase/usecaserole"
	"auth/pkg/logger"
)

type RoleHandler struct {
	usecase usecaserole.IRoleUseCase
	log     logger.ILogger
}

func New(usecase usecaserole.IRoleUseCase) *RoleHandler {
	return &RoleHandler{
		usecase: usecase,
		log:     logger.NewLogger(),
	}
}
