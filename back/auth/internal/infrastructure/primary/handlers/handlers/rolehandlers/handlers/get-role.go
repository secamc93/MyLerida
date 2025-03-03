package handlers

import (
	"auth/internal/domain/role/errors"
	"auth/internal/infrastructure/primary/handlers/handlers/rolehandlers/mappers"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetRole godoc
// @Summary Obtener un rol por ID
// @Description Obtiene la información de un rol específico por su ID
// @Tags roles
// @Accept json
// @Produce json
// @Param id path int true "ID del rol"
// @Success 200 {object} response.RoleResponse "Información del rol"
// @Failure 400 {object} map[string]interface{} "ID de rol inválido"
// @Failure 404 {object} map[string]interface{} "Rol no encontrado"
// @Failure 500 {object} map[string]interface{} "Error interno del servidor"
// @Router /roles/{id} [get]
// @Security ApiKeyAuth
func (h *RoleHandler) GetRole(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		h.log.Error("Error al convertir ID de rol: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de rol inválido"})
		return
	}

	// Llamar al caso de uso
	role, err := h.usecase.GetRoleByID(uint(id))
	if err != nil {
		h.log.Error("Error al obtener rol: %v", err)
		if err == errors.ErrRoleNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Rol no encontrado"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Crear la respuesta utilizando el mapper
	resp := mappers.MapRoleDTOToResponse(role)

	c.JSON(http.StatusOK, resp)
}
