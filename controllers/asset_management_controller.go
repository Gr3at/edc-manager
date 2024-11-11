package controllers

import (
	"edc-proxy/config"
	"edc-proxy/models"
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

	// 2. setup and link connector
	aSConf := AuthStrategyConfig{ClientID: "", ClientSecret: "", TokenURL: ""}
	authType := "JWTAuth"
	authStrategy := NewAuthStrategy(authType, aSConf)

	config := Config{
		ManagementURL: "https://example.com/api/management",
		AuthStrategy:  authStrategy,
	}

	fmt.Println("Get a new API Client instance")
	apiClient, err := NewAPIClient(config, nil)

	if err != nil {
		fmt.Printf("error while creating API client from factory: %v\n", err)
	}

	// 3. propagate json request to the connector
	fmt.Println("Make the Get Assets request")
	byteData, err := apiClient.GetAssets(QueryPayload{
		Type:   "https://w3id.org/edc/v0.0.1/ns/QuerySpec",
		Offset: 0,
		Limit:  20,
	})

	if err != nil {
		fmt.Printf("Error in GetAssets request: %v\n", err)
		// t.Errorf("got status %v but wanted %s", err, "")
		c.AbortWithError(http.StatusBadRequest, err)
	}
	// 4. propagate connector response to end user
	c.JSON(http.StatusOK, gin.H{"response": string(byteData)})
}
