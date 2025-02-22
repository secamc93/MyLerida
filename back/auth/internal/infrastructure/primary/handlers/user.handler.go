package handlers

import (
	"auth/internal/infrastructure/primary/handlers/mappers"
	"auth/internal/infrastructure/primary/handlers/request"
	"auth/internal/infrastructure/primary/handlers/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the input payload
// @Tags users
// @Accept json
// @Produce json
// @Param user body request.UserRequest true "User Request"
// @Success 201 {object} response.BaseResponse
// @Failure 400 {object} response.BaseResponse
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
		c.JSON(http.StatusInternalServerError, response.BaseResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, response.BaseResponse{
		StatusCode: http.StatusCreated,
		Message:    "User created successfully",
	})
}
