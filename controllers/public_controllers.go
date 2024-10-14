package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetServiceHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "Up"})
}
