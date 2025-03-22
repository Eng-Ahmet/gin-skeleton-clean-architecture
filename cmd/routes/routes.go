package routes

import (
	"hwai-api/cmd"
	handler "hwai-api/cmd/Handler"
	"hwai-api/internal/delivery/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRouter sets up the routes for the application
// It returns the gin engine
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Initialize dependencies
	cmd.SetupDependencies()

	// 1️⃣ Add CORS Middleware first
	r.Use(middleware.CORSMiddleware())

	// 2️⃣ Add Error Handling Middleware before executing anything
	r.Use(middleware.ErrorHandlerMiddleware())

	// 3️⃣ Add Authentication Middleware after error handling
	r.Use(middleware.AuthMiddleware())

	// 4️⃣ Load the routes
	// Load the home page
	homePage(r)

	// Create Handlers
	userHandler := handler.SetupUserHandler()
	UserRoutes(r, userHandler)

	// 5️⃣ Add 404 Not Found Middleware after loading all routes
	r.NoRoute(middleware.NotFoundHandlerMiddleware())

	return r
}
