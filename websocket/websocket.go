package websocket

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

// إعداد WebSocket
func SetupWebSocket(app *fiber.App) {
	// إعداد المسار الخاص بـ WebSocket
	app.Get("/ws", websocket.New(func(c *websocket.Conn) {
		// قفل الخيط الحالي مع المعالج
		defer c.Close()

		// تسجيل الاتصال الجديد
		log.Println("New WebSocket connection established")

		// الاستماع للرسائل من العميل
		for {
			msgType, msg, err := c.ReadMessage()
			if err != nil {
				log.Println("Error reading message:", err)
				log.Println("WebSocket connection closed")
				return
			}

			// طباعة الرسالة الواردة
			log.Printf("Received: %s\n", msg)

			// إرسال الرسالة مرة أخرى للعميل (echo)
			if err := c.WriteMessage(msgType, msg); err != nil {
				log.Println("Error sending message:", err)
				return
			}
		}
	}))
}
