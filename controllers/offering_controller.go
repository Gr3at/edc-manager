package controllers

import (
	"edc-proxy/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOffering(c *gin.Context) {
	var input map[string]interface{}

	// Validate input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the service to create resources via third-party API
	err := services.CreateResources(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create offering"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Offering created"})
}
