package handlers

import (
	"net/http"

	"auth/internal/domain/permissions/permissionsdtos"
	"auth/internal/infrastructure/primary/handlers/handlers/permissionshandlers/mappers"

	"github.com/gin-gonic/gin"
)

// Response estándar para el handler GetModules.
type ModulesResponse struct {
	Status  int                         `json:"status"`
	Message string                      `json:"message"`
	Data    []permissionsdtos.ModuleDTO `json:"data"`
}

// GetModules obtiene todos los módulos.
// @Summary Obtener todos los módulos
// @Description Retorna un listado de módulos envuelto en un objeto de respuesta estándar
// @Tags Modulos
// @Produce json
// @Success 200 {object} ModulesResponse
// @Failure 500 {object} map[string]string
// @Router /permissions/modules [get]
func (h *Permissionshandlers) GetModules(c *gin.Context) {
	modules, err := h.usecasepermissions.GetModules()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	moduleResponse := mappers.MapToModuleResponse(modules)

	c.JSON(http.StatusOK, moduleResponse)
}
