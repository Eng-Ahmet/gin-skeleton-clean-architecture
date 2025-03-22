package handler

import (
	"hwai-api/config"
	"hwai-api/internal/delivery/http"
	"hwai-api/internal/repository"
	"hwai-api/internal/usecase"
	"hwai-api/internal/domain"
)

func SetupUserHandler() *http.UserHandler {
	// الاتصال بقاعدة البيانات
	config.ConnectToDatabase()

	// استخدام GORM لإنشاء الجداول تلقائيًا إذا لم تكن موجودة
	config.DB.AutoMigrate(&domain.User{}) // هنا يمكنك إضافة جميع الجداول التي تحتاج إلى إنشائها

	// إنشاء repository و usecase
	userRepo := repository.NewUserRepository(config.DB)
	userUsecase := usecase.NewUserUsecase(*userRepo)
	userHandler := http.NewUserHandler(*userUsecase)

	return userHandler
}
