package router

import (
	"fiper-skeleton-clean-architecture/internal/delivery/middleware"

	"github.com/gofiber/fiber/v2"
)

func InitializeRoutes(app *fiber.App) {
	// Apply middleware
	app.Use(middleware.CorsMiddleware())
	app.Use(middleware.LoggerMiddleware())

	//api := app.Group("/api")

	// User routes
	RegisterUserRoutes(app)
}
