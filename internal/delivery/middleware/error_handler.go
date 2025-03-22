package middleware

import (
	"hwai-api/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrorHandlerMiddleware handles the errors that occur in the application
func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// register the error in the log
				logger.GetErrorLogger().Println("Error occurred in the application with error:", err)
				// return an error response
				c.JSON(http.StatusInternalServerError, gin.H{
					"error":   "Internal Server Error",
					"message": "Something went wrong, please try again later.",
					"code":    http.StatusInternalServerError,
				})
				c.Abort()
			}
		}()

		c.Next()
	}
}
