package handlers

import "auth/internal/application/usecase/usecaseuser"

type UserHandler struct {
	useCase usecaseuser.IUserUseCase
}

func New(useCase usecaseuser.IUserUseCase) *UserHandler {
	return &UserHandler{useCase: useCase}
}
