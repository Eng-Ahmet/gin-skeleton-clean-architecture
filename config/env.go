// config/env.go
package config

import (
	"hwai-api/pkg/logger"
	"os"

	"github.com/joho/godotenv"
)

// if the .env file exists, load it
// if the .env file does not exist, the function will return an error
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		logger.GetErrorLogger().Println("Failed to load .env file")
	}
}

// if the key exists in the environment variables, return its value, otherwise return the default value
func GetEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// set the value of the environment variable
func SetEnv(key string, value string) error {
	err := os.Setenv(key, value)
	if err != nil {
		logger.GetErrorLogger().Println("Failed to set environment variable:", key)
	}
	return err
}
