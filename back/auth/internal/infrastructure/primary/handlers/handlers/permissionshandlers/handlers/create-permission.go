package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"auth/internal/infrastructure/primary/handlers/handlers/permissionshandlers/dtos/request"
	"auth/internal/infrastructure/primary/handlers/handlers/permissionshandlers/mappers"
)

// CreatePermission godoc
// @Summary Crea un nuevo permiso
// @Description Crea un permiso usando los datos provistos en el request
// @Tags permissions
// @Accept json
// @Produce json
// @Param request body request.CreatePermissionRequest true "Datos para crear permiso"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /permissions/create [post]
func (h *Permissionshandlers) CreatePermission(c *gin.Context) {
	var request request.CreatePermissionRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dto := mappers.MapCreatePermissionRequestToDTO(&request)
	_, err := h.usecasepermissions.CreatePermission(dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Creación exitosa"})
}
