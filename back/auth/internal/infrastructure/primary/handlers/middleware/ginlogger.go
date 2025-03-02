package middleware

import (
	"time"

	"auth/pkg/logger"

	"github.com/gin-gonic/gin"
)

// colorStatus devuelve el código de estado formateado según el color ANSI.
func colorStatus(status int) string {
	switch {
	case status >= 200 && status < 300:
		return "\033[32m" // verde
	case status >= 300 && status < 400:
		return "\033[36m" // cyan
	case status >= 400 && status < 500:
		return "\033[33m" // amarillo
	case status >= 500:
		return "\033[31m" // rojo
	default:
		return "\033[0m" // sin color
	}
}

// GinLogger registra los accesos con nuestro logger personalizado.
func GinLogger(l logger.ILogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		latency := time.Since(start)
		statusCode := c.Writer.Status()
		method := c.Request.Method
		path := c.Request.URL.Path
		clientIP := c.ClientIP()

		// Obtener el color ANSI para el status code
		coloredStatus := colorStatus(statusCode)
		// Se imprime el status en color y se resetea al final con "\033[0m"
		l.Info("[GIN] %s %s | %s%d\033[0m | %s | %s", method, path, coloredStatus, statusCode, clientIP, latency)
	}
}
