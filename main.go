package main

import (
	"hwai-api/cmd/routes"
	"hwai-api/config"
)

// main function
func main() {

	// load the environment variables
	config.LoadEnv()

	// read the port from the environment variables
	port := config.GetEnv("PORT", "8001")

	println("")
	println("Server is running on port: " + port + " ...")
	println("API works at: http://localhost:" + port)
	println("Press CTRL + C to stop the server.")
	println("")

	// load the routes
	r := routes.SetupRouter()

	// run the server
	r.Run(":" + port)
}
