package http

import (
	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	// Handle user creation
	return c.JSON(fiber.Map{"message": "User created"})
}

func GetUserByID(c *fiber.Ctx) error {
	// Handle fetching user by ID
	return c.JSON(fiber.Map{"message": "User fetched"})
}

func SayHello(c *fiber.Ctx) error {
	// Handle hello message
	return c.JSON(fiber.Map{"message": "Hello, World!"})
}
