// cmd/routes/user_routes.go
package routes

import (
	"hwai-api/internal/delivery/http"

	"github.com/gin-gonic/gin"
)

// UserRoutes defines routes related to users
func UserRoutes(r *gin.Engine, userHandler *http.UserHandler) {
	// Example of how to use userHandler with routes
	r.POST("/users", userHandler.CreateUser) // Creating a user
	r.GET("/users/:id", userHandler.GetUserByID)             // Getting user by ID
}

// Example route for home page
func homePage(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Welcome to the API"})
	})
}
