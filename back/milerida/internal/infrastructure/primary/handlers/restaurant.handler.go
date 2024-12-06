package handlers

import (
	"my-lerida/internal/application/usecase/restaurant"
	"my-lerida/internal/infrastructure/primary/handlers/mappers"
	"my-lerida/internal/infrastructure/primary/handlers/request"
	"my-lerida/internal/infrastructure/primary/handlers/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RestaurantHandler struct {
	useCase restaurant.IRestaurantUseCase
}

func NewRestaurantHandler(useCase restaurant.IRestaurantUseCase) *RestaurantHandler {
	return &RestaurantHandler{useCase: useCase}
}

func (h *RestaurantHandler) GetRestaurants(c *gin.Context) {
	restaurant := c.Query("restaurant")

	result, err := h.useCase.GetRestaurants(restaurant)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	restaurantDTOs := mappers.ToRestaurantDTOList(result)

	c.JSON(http.StatusOK, response.BaseResponse{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data:       restaurantDTOs,
	})
}

func (h *RestaurantHandler) GetRestaurantByID(c *gin.Context) {
	idParam := c.Param("id")
	restaurantID, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid restaurant ID"})
		return
	}

	restaurant, err := h.useCase.GetRestaurantByID(uint(restaurantID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	restaurantDTO := mappers.ToRestaurantDTO(&restaurant)

	c.JSON(http.StatusOK, response.BaseResponse{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data:       restaurantDTO,
	})
}

func (h *RestaurantHandler) CreateRestaurantHandler(c *gin.Context) {
	var req request.RestaurantRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	restaurant := mappers.ToRestaurantEntitie(&req)

	err := h.useCase.CreateRestaurant(restaurant)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, response.BaseResponse{
		StatusCode: http.StatusCreated,
		Message:    "Restaurant created successfully",
		Data:       nil,
	})

}

func (h *RestaurantHandler) UpdateRestaurantHandler(c *gin.Context) {
	var req request.RestaurantRequest

	idParam := c.Param("id")
	restaurantID, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid restaurant ID"})
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	restaurant := mappers.ToRestaurantEntitie(&req)

	err = h.useCase.UpdateRestaurant(uint(restaurantID), restaurant)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, response.BaseResponse{
		StatusCode: http.StatusCreated,
		Message:    "Restaurant updated successfully",
		Data:       nil,
	})

}

func (ctrl *RestaurantHandler) DeleteRestaurant(c *gin.Context) {
	idParam := c.Param("id")
	restaurantID, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid restaurant ID"})
		return
	}

	err = ctrl.useCase.DeleteRestaurant(uint(restaurantID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.BaseResponse{
		StatusCode: http.StatusAccepted,
		Message:    "Restaurant deleted successfully",
		Data:       nil,
	})
}
