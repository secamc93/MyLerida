package handlers

import (
	"auth/internal/domain/user/usererrors"
	"auth/internal/infrastructure/primary/handlers/handlers/userhandlers/dtos/request"
	"auth/internal/infrastructure/primary/handlers/handlers/userhandlers/dtos/response"
	"auth/internal/infrastructure/primary/handlers/handlers/userhandlers/mappers"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the input payload
// @Tags Users
// @Accept json
// @Produce json
// @Param user body request.UserRequest true "User Request"
// @Success 201 {object} response.BaseResponse
// @Failure 400 {object} response.BaseResponse
// @Failure 409 {object} response.BaseResponse
// @Failure 500 {object} response.BaseResponse
// @Router /users [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	var userRequest request.UserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, response.BaseResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
		return
	}

	userDTO := mappers.MapToUserDTO(userRequest)
	if err := h.useCase.CreateUser(&userDTO); err != nil {
		h.handleCreateUserError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.BaseResponse{
		StatusCode: http.StatusCreated,
		Message:    "Usuario creado exitosamente",
	})
}

func (h *UserHandler) handleCreateUserError(c *gin.Context, err error) {
	var statusCode int
	switch err {
	case usererrors.ErrEmailEmpty, usererrors.ErrNameEmpty, usererrors.ErrPasswordEmpty, usererrors.ErrPasswordInvalid:
		statusCode = http.StatusBadRequest
	case usererrors.ErrEmailAlreadyExists:
		statusCode = http.StatusConflict
	default:
		statusCode = http.StatusInternalServerError
	}
	c.JSON(statusCode, response.BaseResponse{
		StatusCode: statusCode,
		Message:    err.Error(),
	})
}
