package handlers

import (
	"auth/internal/domain/user/errors"
	"auth/internal/infrastructure/primary/handlers/handlers/userhandlers/dtos/response"
	"auth/internal/infrastructure/primary/handlers/handlers/userhandlers/mappers"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUsers godoc
// @Summary Obtiene la lista de usuarios
// @Description Devuelve una lista de todos los usuarios registrados
// @Tags Users
// @Accept json
// @Produce json
// @Success 201 {object} response.BaseResponse
// @Failure 400 {object} response.BaseResponse
// @Failure 404 {object} response.BaseResponse
// @Failure 409 {object} response.BaseResponse
// @Failure 500 {object} response.BaseResponse
// @Router /users [get]
func (h *UserHandler) GetUsers(c *gin.Context) {
	users, err := h.useCase.ListUsers()
	if err != nil {
		if err == errors.ErrNoUsersFound {
			c.JSON(http.StatusNotFound, response.BaseResponse{
				StatusCode: http.StatusNotFound,
				Message:    err.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, response.BaseResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
		return
	}

	// Responder exitosamente con la lista de usuarios
	c.JSON(http.StatusOK, mappers.MapUserResponses(users))
}
