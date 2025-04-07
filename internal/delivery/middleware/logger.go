package middleware

import (
	"fiper-skeleton-clean-architecture/pkg/logger"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func LoggerMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now() // start time measurement

		// Proceed to the next middleware or route handler
		err := c.Next() 

		// Calculate the duration
		duration := time.Since(start)

		// Log the request details
		logMessage := fmt.Sprintf("Request: %s %s | Duration: %s", c.Method(), c.OriginalURL(), duration.String())
		logger.Info(logMessage) // Log the request info

		// Print to console
		fmt.Println(logMessage)

		// إذا كان هناك خطأ، سيتم طباعته هنا
		if err != nil {
			logError := fmt.Sprintf("Error processing request %s %s: %v", c.Method(), c.OriginalURL(), err)
			// سجل الخطأ
			logger.Error(logError)
			// طباعة الخطأ على الكونسول
			fmt.Println(logError)
		}

		return err
	}
}
