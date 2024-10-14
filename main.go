package main

import (
	"edc-proxy/config"
	"edc-proxy/routes"
	"edc-proxy/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the logger
	utils.InitLogger()

	// Initialize the Gin engine
	engine := gin.Default()
	// Set trusted proxies to only those you trust
	// engine.SetTrustedProxies([]string{"192.168.1.1", "127.0.0.1/24"})

	// Load the environment variables or config
	config.LoadConfig()

	// Set up routes
	routes.SetupRoutes(engine)

	// Start the server
	port := config.GetPort()
	utils.Log.Infof("Starting the server on port %s", port)
	if err := engine.Run(":" + port); err != nil {
		utils.Log.Fatalf("Error starting server: %s", err)
	}
}
