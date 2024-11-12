package controllers

import (
	"edc-proxy/pkg/edcclient"
	"edc-proxy/services"
	"edc-proxy/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// func CreateAsset(c *gin.Context) {
// 	var input map[string]interface{}

// 	// Validate input
// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// Call the service to create resources via third-party API
// 	// err := services.CreateResources(input)
// 	// if err != nil {
// 	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create offering"})
// 	// 	return
// 	// }

//		c.JSON(http.StatusOK, gin.H{"status": "Offering created"})
//	}

func CreateAsset(c *gin.Context) {
	// 1. get connector credentials from db
	orgID, exists := c.Get("currentUserOrg")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Your Organization data could not be retrieved"})
		return
	}

	apiClient, err := services.SetupAPIClient(orgID.(string))
	if err != nil {
		utils.Log.Errorf("error creating an apiClient to interact with the edc. error details: (%v)", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 2. Validate input
	var inputPayload edcclient.AnyJSON
	if err := c.ShouldBindJSON(&inputPayload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 3. propagate json request to the connector
	edcResponseBytes, err := apiClient.CreateAsset(inputPayload)
	if err != nil {
		utils.Log.Errorf("Error in CreateAsset request: %v", err)
		c.AbortWithError(http.StatusBadRequest, err)
	}

	c.Data(http.StatusOK, "application/json", edcResponseBytes)
}

func GetAssets(c *gin.Context) {
	// 1. get connector credentials from db
	orgID, exists := c.Get("currentUserOrg")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Your Organization data could not be retrieved"})
		return
	}
	apiClient, err := services.SetupAPIClient(orgID.(string))
	if err != nil {
		utils.Log.Errorf("error creating an apiClient to interact with the edc. error details: (%v)", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 2. Validate input
	var inputQueryPayload edcclient.QueryPayload
	if err := c.ShouldBindJSON(&inputQueryPayload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 3. propagate json request to the connector
	edcResponseBytes, err := apiClient.GetAssets(inputQueryPayload)

	if err != nil {
		// utils.Log.Errorf("Error in GetAssets request: %v", err)
		c.AbortWithError(http.StatusBadRequest, err)
	}

	c.Data(http.StatusOK, "application/json", edcResponseBytes)
}

func DeleteAsset(c *gin.Context) {
	// 1. get connector credentials from db
	orgID, exists := c.Get("currentUserOrg")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Your Organization data could not be retrieved"})
		return
	}

	apiClient, err := services.SetupAPIClient(orgID.(string))
	if err != nil {
		utils.Log.Errorf("error creating an apiClient to interact with the edc. error details: (%v)", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 3. propagate json request to the connector
	assetID := c.Param("assetID")
	edcResponseBytes, err := apiClient.DeleteAsset(assetID)
	_ = edcResponseBytes // dummy line to ignore unwanted variable
	if err != nil {
		utils.Log.Errorf("Error in DeleteAsset request: %v", err)
		c.AbortWithError(http.StatusBadRequest, err)
	}

	c.Status(http.StatusNoContent)
}
