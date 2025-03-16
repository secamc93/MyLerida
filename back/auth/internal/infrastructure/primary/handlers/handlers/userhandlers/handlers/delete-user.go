package handlers

import (
	"auth/internal/domain/user/usererrors"
	"auth/internal/infrastructure/primary/handlers/handlers/userhandlers/dtos/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// DeleteUser godoc
// @Summary Elimina un usuario
// @Description Elimina un usuario por su ID
// @Tags Users
// @Produce json
// @Param id path int true "User ID"
// @Success 201 {object} response.BaseResponse
// @Failure 400 {object} response.BaseResponse
// @Failure 404 {object} response.BaseResponse
// @Failure 409 {object} response.BaseResponse
// @Failure 500 {object} response.BaseResponse
// @Router /users/{id} [delete]
func (h *UserHandler) DeleteUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.BaseResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "ID inválido",
		})
		return
	}

	if err := h.useCase.DeleteUser(uint(id)); err != nil {
		statusCode := http.StatusInternalServerError
		switch err {
		case usererrors.ErrUserNotFound:
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
		Message:    "Usuario eliminado",
	})
}
