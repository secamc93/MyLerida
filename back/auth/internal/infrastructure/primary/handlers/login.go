package handlers

import (
	"auth/internal/domain/user/errors"
	"auth/internal/infrastructure/primary/handlers/dtos/request"
	"auth/internal/infrastructure/primary/handlers/dtos/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login handles the user login process.
// @Summary User login
// @Description Authenticates a user and returns an access token
// @Tags auth
// @Accept json
// @Produce json
// @Param loginRequest body request.LoginRequest true "Login request payload"
// @Success 200 {object} response.LoginResponse
// @Failure 400 {object} response.BaseResponse
// @Failure 401 {object} response.BaseResponse
// @Router /login [post]
func (h *UserHandler) Login(c *gin.Context) {
	var req request.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.BaseResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid request payload",
		})
		return
	}

	token, err := h.useCase.Login(req.Email, req.Password)
	if err != nil {
		h.handleLoginError(c, err)
		return
	}

	resp := response.LoginResponse{AccessToken: token}
	c.JSON(http.StatusOK, resp)
}

func (h *UserHandler) handleLoginError(c *gin.Context, err error) {
	var statusCode int
	switch err {
	case errors.ErrUserNotFound:
		statusCode = http.StatusNotFound
	case errors.ErrInvalidPassword:
		statusCode = http.StatusUnauthorized
	case errors.ErrJWTSecretNotSet:
		statusCode = http.StatusInternalServerError
	default:
		statusCode = http.StatusInternalServerError
	}

	c.JSON(statusCode, response.BaseResponse{
		StatusCode: statusCode,
		Message:    err.Error(),
	})
}
