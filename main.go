package main

import (
	"edc-proxy/config"
	"edc-proxy/routes"
	"edc-proxy/utils"

	"github.com/gin-gonic/gin"
)

func init() {
	// Load the environment variables or config
	config.LoadConfig()
	// Initialize the logger
	utils.InitLogger()
	utils.Log.Info("Connecting DB...")
	config.ConnectDB()
	utils.Log.Info("DB Connected.")
}

func main() {
	// Initialize the Gin engine
	engine := gin.Default()
	// Set trusted proxies to only those you trust
	// engine.SetTrustedProxies([]string{"192.168.1.1", "127.0.0.1/24"})

	// Set up routes
	routes.SetupRoutes(engine)

	// Start the server
	host := config.GetHost()
	port := config.GetPort()
	utils.Log.Infof("Starting the server on port %s", port)
	if err := engine.Run(host + ":" + port); err != nil {
		utils.Log.Fatalf("Error starting server: %s", err)
	}
}
