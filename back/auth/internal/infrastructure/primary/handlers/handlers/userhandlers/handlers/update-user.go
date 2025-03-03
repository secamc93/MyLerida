package handlers

import (
	"auth/internal/infrastructure/primary/handlers/handlers/userhandlers/dtos/request"
	"auth/internal/infrastructure/primary/handlers/handlers/userhandlers/dtos/response"
	"auth/internal/infrastructure/primary/handlers/handlers/userhandlers/mappers"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UpdateUser godoc
// @Summary Actualiza un usuario
// @Description Actualiza los datos de un usuario por su ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body request.UserUpdateRequest true "User update payload"
// @Success 200 {object} response.BaseResponse
// @Failure 400 {object} response.BaseResponse
// @Failure 404 {object} response.BaseResponse
// @Failure 409 {object} response.BaseResponse
// @Failure 500 {object} response.BaseResponse
// @Router /users/{id} [put]
func (h *UserHandler) UpdateUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.BaseResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "ID inválido",
		})
		return
	}

	var req request.UserUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.BaseResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Payload inválido",
		})
		return
	}

	userDTO := mappers.ToUserUpdateDTO(req)

	if err := h.useCase.UpdateUser(uint(id), &userDTO); err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "usuario no encontrado" {
			statusCode = http.StatusNotFound
		}
		c.JSON(statusCode, response.BaseResponse{
			StatusCode: statusCode,
			Message:    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.BaseResponse{
		StatusCode: http.StatusOK,
		Message:    "Usuario actualizado",
	})
}
