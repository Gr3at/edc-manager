package controllers

import (
	"edc-proxy/config"
	"edc-proxy/models"
	"edc-proxy/pkg/edcclient"
	"edc-proxy/utils"
	"fmt"
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
// 	err := services.CreateResources(input)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create offering"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"status": "Offering created"})
// }

func GetAssets(c *gin.Context) {
	// 1. get connector credentials from db
	orgID, exists := c.Get("currentUserOrg")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Organization ID not found in context"})
		return
	}

	var orgConnector models.Connector
	result := config.DB.Where("org_id = ?", orgID).First(&orgConnector)

	if result.RowsAffected == 0 {
		utils.Log.Warnf("No connector credentials found for the requested org: %s", orgID)
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("You need to link your organization (%s) Connector first.", orgID)})
		return
	}

	if result.Error != nil {
		utils.Log.Errorf("Error fetching the organization's connector: %v", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to use your connector."})
		return
	}

	var inputQueryPayload edcclient.QueryPayload

	// Validate input
	if err := c.ShouldBindJSON(&inputQueryPayload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 2. setup and link connector
	panic("Implement logic to read credentials from db and split to id and secret. consider adding token url in the db record")
	aSConf := edcclient.AuthStrategyConfig{ClientID: "", ClientSecret: "", TokenURL: ""}
	authStrategy := edcclient.NewAuthStrategy(string(orgConnector.CredentialsType), aSConf)

	config := edcclient.Config{
		ManagementURL: orgConnector.APIUrl,
		AuthStrategy:  authStrategy,
	}

	utils.Log.Info("Get a new API Client instance")
	apiClient, err := edcclient.NewAPIClient(config, nil)

	if err != nil {
		utils.Log.Errorf("error while creating API client from factory: %v\n", err)
	}

	// 3. propagate json request to the connector
	utils.Log.Info("Make the Get Assets request")
	edcResponseBytes, err := apiClient.GetAssets(inputQueryPayload)

	if err != nil {
		utils.Log.Errorf("Error in GetAssets request: %v\n", err)
		c.AbortWithError(http.StatusBadRequest, err)
	}

	c.Data(http.StatusOK, "application/json", edcResponseBytes)
	// var edcResponsePayload []map[string]interface{}

	// // Validate input

	// if err := json.Unmarshal(edcResponseBytes, &edcResponsePayload); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }
	// // 4. propagate connector response to end user
	// c.JSON(http.StatusOK, edcResponsePayload)
}
