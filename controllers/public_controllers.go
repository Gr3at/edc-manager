package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetServiceStatus(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "Up"})
}
