package config

import (
	"fmt"
	"hwai-api/pkg/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() {
	// Replace with your MySQL credentials
	dsn := "root:@tcp(localhost:3306)/hwai-go?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.GetErrorLogger().Println("Failed to connect to the MySQL database")
		logger.GetInfoLogger().Println(err.Error())
		return
	}

	// Test the connection (optional but good practice)
	sqlDB, err := DB.DB()
	if err != nil {
		logger.GetErrorLogger().Println("Failed to get the database object for ping")
		return
	}

	// Ping the database to ensure the connection is alive
	err = sqlDB.Ping()
	if err != nil {
		logger.GetErrorLogger().Println("Failed to ping the MySQL database")
		return
	}

	fmt.Println("Successfully connected to the MySQL database!")
}
