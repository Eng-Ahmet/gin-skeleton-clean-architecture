package cmd

import (
	"hwai-api/config"
)

func SetupDependencies() {
	config.ConnectToDatabase()
}
