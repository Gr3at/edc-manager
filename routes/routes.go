package routes

import (
	"edc-proxy/controllers"
	"edc-proxy/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Public routes - None
	router.GET("/api/v1/status", controllers.GetServiceStatus)

	// Protected routes (require JWT)
	protected := router.Group("/api/v1/connector")
	protected.Use(middleware.JWTAuthMiddleware())
	{
		protected.POST("/", controllers.CreateConnector)
		protected.GET("/", controllers.GetOrgConnector)
		protected.PUT("/", controllers.UpdateConnector)
		protected.DELETE("/", controllers.DeleteConnector)
	}
}
