package routes

import (
	"edc-proxy/controllers"
	"edc-proxy/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Public routes - None
	router.GET("/health", controllers.GetServiceHealth)

	// Protected routes (require JWT)
	protected := router.Group("/")
	protected.Use(middleware.JWTAuthMiddleware())
	{
		router.POST("/connectors", controllers.AddConnector)
		router.GET("/connectors", controllers.GetConnectors)
		protected.POST("/create-offering", controllers.CreateOffering)
	}
}
