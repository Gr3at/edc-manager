package controllers

import (
	"edc-proxy/pkg/edcclient"
	"edc-proxy/services"
	"edc-proxy/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getAPIClient(c *gin.Context) (*edcclient.APIClient, error) {
	orgID, exists := c.Get("currentUserOrg")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Your Organization data could not be retrieved"})
		return nil, fmt.Errorf("organization ID not found")
	}

	apiClient, err := services.SetupAPIClient(orgID.(string))
	if err != nil {
		utils.Log.Errorf("error creating an apiClient to interact with the edc. error details: (%v)", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil, err
	}

	return apiClient, nil
}

func handleErrorResponse(c *gin.Context, err error, statusCode int) {
	utils.Log.Errorf("Error: %v", err)
	c.AbortWithError(statusCode, err)
}

func handleSuccessResponse(c *gin.Context, data []byte) {
	c.Data(http.StatusOK, "application/json", data)
}

func CreateAsset(c *gin.Context) {
	apiClient, err := getAPIClient(c)
	if err != nil {
		return
	}

	var inputPayload edcclient.AnyJSON
	if err := c.ShouldBindJSON(&inputPayload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	edcResponseBytes, err := apiClient.CreateAsset(inputPayload)
	if err != nil {
		handleErrorResponse(c, err, http.StatusBadRequest)
		return
	}

	handleSuccessResponse(c, edcResponseBytes)
}

func UpdateAsset(c *gin.Context) {
	apiClient, err := getAPIClient(c)
	if err != nil {
		return
	}

	var inputPayload edcclient.AnyJSON
	if err := c.ShouldBindJSON(&inputPayload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	edcResponseBytes, err := apiClient.UpdateAsset(inputPayload)
	if err != nil {
		handleErrorResponse(c, err, http.StatusBadRequest)
		return
	}

	handleSuccessResponse(c, edcResponseBytes)
}

func GetAssets(c *gin.Context) {
	apiClient, err := getAPIClient(c)
	if err != nil {
		return
	}

	var inputQueryPayload edcclient.QueryPayload
	if err := c.ShouldBindJSON(&inputQueryPayload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	edcResponseBytes, err := apiClient.GetAssets(inputQueryPayload)
	if err != nil {
		handleErrorResponse(c, err, http.StatusBadRequest)
		return
	}

	handleSuccessResponse(c, edcResponseBytes)
}

func DeleteAsset(c *gin.Context) {
	apiClient, err := getAPIClient(c)
	if err != nil {
		return
	}

	assetID := c.Param("assetID")
	edcResponseBytes, err := apiClient.DeleteAsset(assetID)
	_ = edcResponseBytes
	if err != nil {
		handleErrorResponse(c, err, http.StatusBadRequest)
		return
	}

	c.Status(http.StatusNoContent)
}
