package handlers

import (
	"auth/internal/infrastructure/primary/handlers/handlers/rolehandlers/dtos/request"
	"auth/internal/infrastructure/primary/handlers/handlers/rolehandlers/mappers"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateRole godoc
// @Summary Crear un nuevo rol
// @Description Crea un nuevo rol en el sistema
// @Tags roles
// @Accept json
// @Produce json
// @Param role body request.CreateRoleRequest true "Datos del rol"
// @Success 201 {object} response.RoleResponse "Rol creado"
// @Failure 400 {object} map[string]interface{} "Error en los datos de solicitud"
// @Failure 409 {object} map[string]interface{} "Ya existe un rol con ese nombre"
// @Failure 500 {object} map[string]interface{} "Error interno del servidor"
// @Router /roles [post]
// @Security ApiKeyAuth
func (h *RoleHandler) CreateRole(c *gin.Context) {
	var req request.CreateRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Error("Error al vincular JSON de solicitud: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos de solicitud inválidos"})
		return
	}

	roleDTO := mappers.MapCreateRequestToRoleDTO(&req)
	createdRole, err := h.usecase.CreateRole(roleDTO)
	if err != nil {
		h.log.Error("Error al crear rol: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := mappers.MapRoleDTOToResponse(createdRole)

	c.JSON(http.StatusCreated, resp)
}
