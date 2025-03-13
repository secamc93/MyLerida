package handlers

import (
	"auth/internal/domain/role/roleerrors"
	"auth/internal/infrastructure/primary/handlers/handlers/rolehandlers/dtos/request"
	"auth/internal/infrastructure/primary/handlers/handlers/rolehandlers/mappers"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UpdateRole godoc
// @Summary Actualizar un rol
// @Description Actualiza la información de un rol existente
// @Tags roles
// @Accept json
// @Produce json
// @Param id path int true "ID del rol"
// @Param role body request.UpdateRoleRequest true "Datos actualizados del rol"
// @Success 200 {object} response.RoleResponse "Rol actualizado"
// @Failure 400 {object} map[string]interface{} "Datos de solicitud inválidos"
// @Failure 404 {object} map[string]interface{} "Rol no encontrado"
// @Failure 409 {object} map[string]interface{} "Ya existe un rol con ese nombre"
// @Failure 500 {object} map[string]interface{} "Error interno del servidor"
// @Router /roles/{id} [put]
// @Security ApiKeyAuth
func (h *RoleHandler) UpdateRole(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		h.log.Error("Error al convertir ID de rol: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de rol inválido"})
		return
	}

	var req request.UpdateRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Error("Error al vincular JSON de solicitud: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos de solicitud inválidos"})
		return
	}

	// Convertir la solicitud al DTO del dominio usando el mapper
	roleDTO := mappers.MapUpdateRequestToRoleDTO(&req, uint(id))

	// Llamar al caso de uso
	updatedRole, err := h.usecase.UpdateRole(uint(id), roleDTO)
	if err != nil {
		h.log.Error("Error al actualizar rol: %v", err)
		if err == roleerrors.ErrRoleNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Rol no encontrado"})
			return
		}
		if err == roleerrors.ErrRoleAlreadyExists {
			c.JSON(http.StatusConflict, gin.H{"error": "Ya existe un rol con ese nombre"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Crear la respuesta utilizando el mapper
	resp := mappers.MapRoleDTOToResponse(updatedRole)

	c.JSON(http.StatusOK, resp)
}
