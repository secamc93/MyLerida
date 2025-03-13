package handlers

import (
	"auth/internal/domain/role/roleerrors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// DeleteRole godoc
// @Summary Eliminar un rol
// @Description Elimina un rol específico por su ID
// @Tags roles
// @Accept json
// @Produce json
// @Param id path int true "ID del rol"
// @Success 200 {object} map[string]interface{} "Mensaje de éxito"
// @Failure 400 {object} map[string]interface{} "ID de rol inválido"
// @Failure 404 {object} map[string]interface{} "Rol no encontrado"
// @Failure 500 {object} map[string]interface{} "Error interno del servidor"
// @Router /roles/{id} [delete]
// @Security ApiKeyAuth
func (h *RoleHandler) DeleteRole(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		h.log.Error("Error al convertir ID de rol: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de rol inválido"})
		return
	}

	// Llamar al caso de uso
	err = h.usecase.DeleteRole(uint(id))
	if err != nil {
		h.log.Error("Error al eliminar rol: %v", err)
		if err == roleerrors.ErrRoleNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Rol no encontrado"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Rol eliminado correctamente"})
}
