package controllers

import (
	"edc-proxy/config"
	"edc-proxy/models"
	"edc-proxy/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type connectorInput struct {
	APIUrl string `json:"api_url" binding:"required;min=10"`
	APIKey string `json:"api_key" binding:"required;min=2"`
	// SubID                  string `json:"sub_id" binding:"required"`
	// OrgID                  string `json:"org_id" binding:"required"`
	// AvailableToAllOrgUsers bool `json:"available_to_all_org_users"`
}

type connectorOutput struct {
	ID     uuid.UUID `json:"id"`
	APIUrl string    `json:"api_url"`
}

func AddConnector(c *gin.Context) {
	var input connectorInput

	// Validate JSON input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	subID, exists := c.Get("currentUserSub")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Sub ID not found in context"})
		return
	}
	orgID, exists := c.Get("currentUserOrg")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Organization ID not found in context"})
		return
	}
	utils.Log.Debugf("AddConnector context: %s - %s", subID, orgID)

	secretKey := config.GetSecret()
	encryptedAPIKey, err := utils.EncryptAPIKey(input.APIKey, secretKey)
	if err != nil {
		utils.Log.Errorf("Error adding new record. ApiKey error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to create record. ApiKey error."})
		return
	}

	record := models.Connector{
		APIUrl:         input.APIUrl,
		APIKey:         encryptedAPIKey, // Keep this secret
		UpdatedBySubID: subID.(string),
		OrgID:          orgID.(string),
		// AvailableToAllOrgUsers: input.AvailableToAllOrgUsers,
	}

	// create db record
	result := config.DB.Create(&record)
	if result.Error != nil {
		utils.Log.Errorf("Error adding new record: %v", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Unable to create record",
			"details": result.Error.Error(),
		})
		return
	}
	utils.Log.Infof("New record committed with id: %s", record.ID)

	c.JSON(http.StatusCreated, record)
}

func GetOrgConnector(c *gin.Context) {
	var orgConnector models.Connector

	orgID, exists := c.Get("currentUserOrg")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Organization ID not found in context"})
		return
	}

	result := config.DB.Where("org_id = ?", orgID).First(&orgConnector)

	if result.Error != nil {
		utils.Log.Errorf("Error fetching the organization's connector: %v", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to retrieve the organization's connector"})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusOK, connectorOutput{})
	} else {
		output := connectorOutput{
			ID:     orgConnector.ID,
			APIUrl: orgConnector.APIUrl,
		}
		c.JSON(http.StatusOK, output)
	}
}

func UpdateConnector(c *gin.Context) {
	var input connectorInput

	// Validate JSON input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// start - retrieve db connector record
	var orgConnector models.Connector

	orgID, exists := c.Get("currentUserOrg")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Organization ID not found in context"})
		return
	}

	result := config.DB.Where("org_id = ?", orgID).First(&orgConnector)

	if (result.Error != nil) || (result.RowsAffected == 0) {
		utils.Log.Errorf("Cannot locate and retrieve the organization's connector: %v", result.Error)
		c.JSON(http.StatusNotFound, gin.H{"error": "Cannot update the organization's connector"})
		return
	}
	// end - retrieve db connector record

	// create db record
	orgConnector.APIUrl = input.APIUrl
	orgConnector.APIKey = input.APIKey
	saveResult := config.DB.Save(&orgConnector)

	if saveResult.Error != nil {
		utils.Log.Errorf("Error updating record: %v", saveResult.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Unable to update record",
			"details": saveResult.Error.Error(),
		})
		return
	}
	utils.Log.Infof("Record update committed with id: %s", orgConnector.ID)

	output := connectorOutput{
		ID:     orgConnector.ID,
		APIUrl: orgConnector.APIUrl,
	}

	c.JSON(http.StatusOK, output)
}

func DeleteConnector(c *gin.Context) {
	// Retrieve organization ID from JWT claims (set in middleware)
	orgID, exists := c.Get("currentUserOrg")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Organization ID not found"})
		return
	}

	var orgConnector models.Connector
	result := config.DB.Where("org_id = ?", orgID).First(&orgConnector)

	if result.Error != nil {
		utils.Log.Errorf("Error retrieving connector: %v", result.Error)
		c.JSON(http.StatusNotFound, gin.H{"error": "Connector not found"})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Connector not found for the organization"})
		return
	}

	// Delete the connector
	deleteResult := config.DB.Delete(&orgConnector)
	if deleteResult.Error != nil {
		utils.Log.Errorf("Error deleting connector: %v", deleteResult.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to delete the connector"})
		return
	}

	utils.Log.Infof("Connector deleted successfully for OrgID: %s", orgID)
	c.JSON(http.StatusNoContent, gin.H{"message": "Connector deleted successfully"})
}
