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
	setupPolicyRoutes(protected)
	setupContractDefinitionRoutes(protected)
	setupContractNegotiationRoutes(protected)
	setupContractAgreementRoutes(protected)
	setupCatalogueRoutes(protected)
}

func setupConnectorRoutes(group *gin.RouterGroup) {
	connectorGroup := group.Group("/connector")
	connectorGroup.POST("/", controllers.CreateConnector)
	connectorGroup.GET("/", controllers.GetOrgConnector)
	connectorGroup.PUT("/", controllers.UpdateConnector)
	connectorGroup.DELETE("/", controllers.DeleteConnector)
}

func setupAssetsRoutes(group *gin.RouterGroup) {
	assetsGroup := group.Group("/assets")
	assetsGroup.POST("/", controllers.CreateAsset)           // Create asset
	assetsGroup.PUT("/", controllers.UpdateAsset)            // Update asset
	assetsGroup.POST("/request", controllers.GetAssets)      // List assets
	assetsGroup.DELETE("/:assetID", controllers.DeleteAsset) // Delete asset
}

func setupPolicyRoutes(group *gin.RouterGroup) {
	policyGroup := group.Group("/policydefinitions")
	policyGroup.POST("/", controllers.CreatePolicy)            // Create policy
	policyGroup.POST("/request", controllers.GetPolicies)      // List policies
	policyGroup.DELETE("/:policyID", controllers.DeletePolicy) // Delete policy
}

func setupContractDefinitionRoutes(group *gin.RouterGroup) {
	contractDefinitionsGroup := group.Group("/contractdefinitions")
	contractDefinitionsGroup.POST("/", controllers.CreateContractDefinition)
	contractDefinitionsGroup.POST("/request", controllers.GetContractDefinitions)
	contractDefinitionsGroup.DELETE("/:contractDefinitionID", controllers.DeleteContractDefinition)
}

func setupContractNegotiationRoutes(group *gin.RouterGroup) {
	contractNegotiationsGroup := group.Group("/contractnegotiations")
	contractNegotiationsGroup.POST("/", controllers.StartContractNegotiation)
	contractNegotiationsGroup.POST("/request", controllers.GetContractNegotiations)
	contractNegotiationsGroup.POST("/:id/cancel", controllers.CancelContractNegotiation)
	contractNegotiationsGroup.POST("/:id/decline", controllers.DeclineContractNegotiation)
}

func setupContractAgreementRoutes(group *gin.RouterGroup) {
	contractAgreementsGroup := group.Group("/contractagreements")
	contractAgreementsGroup.POST("/request", controllers.GetContractAgreements)
}

func setupCatalogueRoutes(group *gin.RouterGroup) {
	catalogueGroup := group.Group("/catalog")
	catalogueGroup.POST("/request", controllers.RequestCatalogue)
}
