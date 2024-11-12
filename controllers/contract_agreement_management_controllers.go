package controllers

import (
	"edc-proxy/pkg/edcclient"
	"edc-proxy/services"
	"edc-proxy/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetContractAgreements(c *gin.Context) {
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
	edcResponseBytes, err := apiClient.GetContractAgreements(inputQueryPayload)

	if err != nil {
		utils.Log.Errorf("Error in GetContractAgreements request: %v", err)
		c.AbortWithError(http.StatusBadRequest, err)
	}

	c.Data(http.StatusOK, "application/json", edcResponseBytes)
}
