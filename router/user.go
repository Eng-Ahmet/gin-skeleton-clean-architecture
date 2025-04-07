package router

import (
	"fiper-skeleton-clean-architecture/internal/delivery/http"

	"github.com/gofiber/fiber/v2"
)

func RegisterUserRoutes(router fiber.Router) {
	router.Get("/", http.SayHello)
	router.Get("/user/:id", http.GetUserByID)
	router.Post("/", http.CreateUser)
}
