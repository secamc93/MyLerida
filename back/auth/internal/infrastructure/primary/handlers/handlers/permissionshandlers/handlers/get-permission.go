package handlers

import (
	"auth/internal/infrastructure/primary/handlers/handlers/permissionshandlers/mappers"
	"auth/pkg/globalerrors"
	"auth/pkg/utils"

	"github.com/gin-gonic/gin"
)

// GetPermission retrieves the permissions for a specific business and user.
//
// @Summary      Retrieve permissions for a business and user
// @Description  Extracts the business ID and user ID from the request headers, then fetches associated permissions.
// @Tags         Permissions
// @Accept       json
// @Produce      json
// @Param        businesses_id  header    string  true  "Business ID"
// @Param        user_id        header    string  true  "User ID"
// @Success      200  {array}   response.PermissionResponse    "List of permissions"
// @Failure      400  {object}  globalerrors.BaseResponse    "Header conversion error or missing required header"
// @Failure      500  {object}  globalerrors.BaseResponse    "Error retrieving permissions"
// @Router       /permission [get]
func (h *Permissionshandlers) GetPermission(c *gin.Context) {
	businessID, err := utils.HeaderToUint(c, "businesses_id")
	if err != nil {
		c.JSON(400, globalerrors.BaseResponse{
			StatusCode: 400,
			Message:    err.Error(),
		})
		return
	}

	userID, err := utils.HeaderToUint(c, "user_id")
	if err != nil {
		c.JSON(400, globalerrors.BaseResponse{
			StatusCode: 400,
			Message:    err.Error(),
		})
		return
	}

	permissionDTOs, err := h.usecasepermissions.GetPermissionByBussinesAndUser(c.Request.Context(), businessID, userID)
	if err != nil {
		c.JSON(500, globalerrors.BaseResponse{
			StatusCode: 500,
			Message:    "error al obtener los permisos",
		})
		return
	}

	responseDTOs := mappers.MapToPermissionResponseList(*permissionDTOs)
	c.JSON(200, responseDTOs)
}
