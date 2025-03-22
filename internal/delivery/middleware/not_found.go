package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// NotFoundHandlerMiddleware يتعامل مع المسارات غير الموجودة
func NotFoundHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Not Found",
			"message": "The requested resource was not found.",
			"code":    http.StatusNotFound,
		})
	}
}
