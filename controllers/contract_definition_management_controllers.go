package controllers

import (
	"edc-proxy/pkg/edcclient"
	"edc-proxy/services"
	"edc-proxy/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateContractDefinition(c *gin.Context) {
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
	var inputPayload edcclient.ContractDefinition
	if err := c.ShouldBindJSON(&inputPayload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 3. propagate json request to the connector
	edcResponseBytes, err := apiClient.CreateContractDefinition(inputPayload)
	if err != nil {
		utils.Log.Errorf("Error in CreateContractDefinition request: %v", err)
		c.AbortWithError(http.StatusBadRequest, err)
	}

	c.Data(http.StatusOK, "application/json", edcResponseBytes)
}

func GetContractDefinitions(c *gin.Context) {
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
	edcResponseBytes, err := apiClient.GetContractDefinitions(inputQueryPayload)

	if err != nil {
		utils.Log.Errorf("Error in GetContractDefinitions request: %v", err)
		c.AbortWithError(http.StatusBadRequest, err)
	}

	c.Data(http.StatusOK, "application/json", edcResponseBytes)
}

func DeleteContractDefinition(c *gin.Context) {
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
	contractDefinitionID := c.Param("contractDefinitionID")
	edcResponseBytes, err := apiClient.DeleteContractDefinition(contractDefinitionID)
	_ = edcResponseBytes // dummy line to ignore unwanted variable
	if err != nil {
		utils.Log.Errorf("Error in DeleteContractDefinition request: %v", err)
		c.AbortWithError(http.StatusBadRequest, err)
	}

	c.Status(http.StatusNoContent)
}
