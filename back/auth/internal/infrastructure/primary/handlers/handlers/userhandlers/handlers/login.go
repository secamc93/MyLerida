package handlers

import (
	"auth/internal/domain/user/usererrors"
	"auth/internal/infrastructure/primary/handlers/handlers/userhandlers/dtos/request"
	"auth/internal/infrastructure/primary/handlers/handlers/userhandlers/dtos/response"
	"auth/internal/infrastructure/primary/handlers/handlers/userhandlers/mappers"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login handles the user login process.
// @Summary Login
// @Description Authenticates a user and returns an access token
// @Tags Login
// @Accept json
// @Produce json
// @Param loginRequest body request.LoginRequest true "Login request payload"
// @Success 200 {object} response.LoginResponse
// @Failure 400 {object} response.BaseResponse
// @Failure 401 {object} response.BaseResponse
// @Router /users/login [post]
func (h *UserHandler) Login(c *gin.Context) {
	var req request.LoginRequest
	ctx := c.Request.Context()

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.BaseResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid request payload",
		})
		return
	}

	loginRespDTO, err := h.useCase.Login(ctx, req.Email, req.Password)
	if err != nil {
		h.handleLoginError(c, err)
		return
	}

	resp := mappers.MapLoginResponse(loginRespDTO)
	c.JSON(http.StatusOK, resp)
}

func (h *UserHandler) handleLoginError(c *gin.Context, err error) {
	var statusCode int
	switch err {
	case usererrors.ErrUserNotFound:
		statusCode = http.StatusNotFound
	case usererrors.ErrInvalidPassword:
		statusCode = http.StatusUnauthorized
	case usererrors.ErrJWTSecretNotSet:
		statusCode = http.StatusInternalServerError
	default:
		statusCode = http.StatusInternalServerError
	}

	c.JSON(statusCode, response.BaseResponse{
		StatusCode: statusCode,
		Message:    err.Error(),
	})
}
