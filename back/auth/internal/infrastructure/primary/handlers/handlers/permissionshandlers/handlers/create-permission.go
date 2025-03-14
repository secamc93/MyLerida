package handlers

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
