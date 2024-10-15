package controllers

import (
	"edc-proxy/config"
	"edc-proxy/models"
	"edc-proxy/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddConnector(c *gin.Context) {
	var input models.ConnectorInput

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

	record := models.Connector{
		APIUrl:                 input.APIUrl,
		APIKey:                 input.APIKey, // Keep this secret
		SubID:                  subID.(string),
		OrgID:                  orgID.(string),
		AvailableToAllOrgUsers: input.AvailableToAllOrgUsers,
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

func GetConnectors(c *gin.Context) {
	var connectors []models.Connector

	orgID, exists := c.Get("currentUserOrg")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Organization ID not found in context"})
		return
	}

	// result := config.DB.Find(&connectors)
	result := config.DB.Where("org_id = ?", orgID).Find(&connectors)

	if result.Error != nil {
		utils.Log.Errorf("Error fetching connectors: %v", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch connectors"})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusOK, []models.Connector{})
	} else {
		c.JSON(http.StatusOK, connectors)
	}

}
