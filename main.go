package main

import (
	"fiper-skeleton-clean-architecture/router"
	"fiper-skeleton-clean-architecture/pkg/logger"
	"fiper-skeleton-clean-architecture/websocket"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// تهيئة السجل
	logger.InitLogger("log/info.log", "log/errors.log")

	// إنشاء تطبيق Fiber بدون تفعيل Prefork
	app := fiber.New(fiber.Config{
		//Prefork: true, // إلغاء Prefork
	})

	// إضافة Middleware CORS
	app.Use(cors.New())

	// إعداد مسارات API
	router.InitializeRoutes(app)

	// إعداد WebSocket
	websocket.SetupWebSocket(app)

	// بدء السيرفر على المنفذ 8080
	if err := app.Listen(":8080"); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
