package handlers

import (
	"auth/internal/infrastructure/primary/handlers/handlers/rolehandlers/mappers"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ListRoles godoc
// @Summary Listar todos los roles
// @Description Obtiene una lista de todos los roles disponibles en el sistema
// @Tags roles
// @Accept json
// @Produce json
// @Success 200 {object} response.RoleListResponse "Lista de roles"
// @Failure 500 {object} map[string]interface{} "Error interno del servidor"
// @Router /roles [get]
// @Security ApiKeyAuth
func (h *RoleHandler) ListRoles(c *gin.Context) {
	// Llamar al caso de uso
	roles, err := h.usecase.ListRoles()
	if err != nil {
		h.log.Error("Error al listar roles: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Crear la respuesta utilizando el mapper
	resp := mappers.MapRoleDTOsToResponseList(roles)

	c.JSON(http.StatusOK, resp)
}
