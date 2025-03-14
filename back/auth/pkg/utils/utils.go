package utils

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func HeaderToUint(c *gin.Context, headerName string) (uint, error) {
	value := c.GetHeader(headerName)
	if value == "" {
		return 0, fmt.Errorf("%s es requerido en el header", headerName)
	}
	parsed, err := strconv.ParseUint(value, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(parsed), nil
}
