package routes

import (
	"edc-proxy/controllers"
	"edc-proxy/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Public routes - None
	router.GET("/status", controllers.GetServiceStatus)

	// Protected routes (require JWT)
	protected := router.Group("/")
	protected.Use(middleware.JWTAuthMiddleware())
	{
		protected.POST("/connector", controllers.AddConnector)
		protected.GET("/connector", controllers.GetOrgConnector)
		protected.PUT("/connector", controllers.UpdateConnector)
		protected.DELETE("/connector", controllers.DeleteConnector)
	}
}
