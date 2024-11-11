package routes

import (
	"edc-proxy/controllers"
	"edc-proxy/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Public routes - None
	router.GET("/api/v1/status", controllers.GetServiceStatus)

	// Protected connector credentials routes (require JWT)
	protected := router.Group("/api/v1")
	protected.Use(middleware.JWTAuthMiddleware())

	setupConnectorRoutes(protected)
	setupAssetsRoutes(protected)
}

func setupConnectorRoutes(group *gin.RouterGroup) {
	connector := group.Group("/connector")
	connector.POST("/", controllers.CreateConnector)
	connector.GET("/", controllers.GetOrgConnector)
	connector.PUT("/", controllers.UpdateConnector)
	connector.DELETE("/", controllers.DeleteConnector)
}

func setupAssetsRoutes(group *gin.RouterGroup) {
	assets := group.Group("/assets")
	// assets.POST("/", controllers.CreateConnector)        // Create asset
	assets.POST("/request", controllers.GetAssets) // List assets
	// assets.DELETE("/", controllers.DeleteConnector)      // Delete asset
}
