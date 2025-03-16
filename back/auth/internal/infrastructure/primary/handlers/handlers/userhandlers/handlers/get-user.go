package handlers

import (
	"auth/internal/domain/user/usererrors"
	"auth/internal/infrastructure/primary/handlers/handlers/userhandlers/dtos/response"
	"auth/internal/infrastructure/primary/handlers/handlers/userhandlers/mappers"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetUser godoc
// @Summary Obtiene un usuario por su ID
// @Description Devuelve un usuario específico basado en su ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} response.UserResponse
// @Failure 400 {object} response.BaseResponse
// @Failure 404 {object} response.BaseResponse
// @Failure 500 {object} response.BaseResponse
// @Router /users/{id} [get]
func (h *UserHandler) GetUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.BaseResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "ID inválido",
		})
		return
	}

	user, err := h.useCase.GetUserByID(uint(id))
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err == usererrors.ErrUserNotFound {
			statusCode = http.StatusNotFound
		}
		c.JSON(statusCode, response.BaseResponse{
			StatusCode: statusCode,
			Message:    err.Error(),
		})
		return
	}

	// Convertir el UserDTO a UserResponse
	userResponse := mappers.MapUserResponse(*user)
	c.JSON(http.StatusOK, userResponse)
}
