package controllers

import (
	"edc-proxy/models"
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

	// Create record (in a real app, you would also interact with a DB here)
	record := models.Connector(input)
	// record := models.Connector{
	// 	APIUrl: input.APIUrl,
	// 	// APIKey: "input.APIKey", // Keep this secret
	// 	APIKey: input.APIKey, // Keep this secret
	// 	// SubID:                  input.SubID,
	// 	AvailableToAllOrgUsers: input.AvailableToAllOrgUsers,
	// }

	// Save to DB (simulate for now)
	c.JSON(http.StatusCreated, gin.H{"data": record})
}

func GetConnectors(c *gin.Context) {
	// Simulate fetching 2 records from the database (non-secret info only)
	record1 := models.Connector{
		APIUrl:                 "https://example2.com/api",
		AvailableToAllOrgUsers: true,
	}
	record2 := models.Connector{
		APIUrl:                 "https://example2.com/api",
		AvailableToAllOrgUsers: false,
	}

	records := []models.Connector{record1, record2}

	c.JSON(http.StatusOK, gin.H{"connectors": records})
}
