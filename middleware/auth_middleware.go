package middleware

import (
	"edc-proxy/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the JWT token from the request header
		tokenString := c.GetHeader("Authorization")

		// Log the request with the token (if present)
		utils.Log.WithFields(logrus.Fields{
			"token": tokenString,
		}).Info("Validating JWT token")

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token not provided"})
			c.Abort()
			return
		}

		// Validate the token
		_, err := utils.ValidateJWT(tokenString)
		if err != nil {
			// Log the error
			utils.Log.WithFields(logrus.Fields{
				"error": err.Error(),
				"token": tokenString,
			}).Error("JWT validation failed")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Log successful validation
		utils.Log.WithFields(logrus.Fields{
			"token": tokenString,
		}).Info("JWT token validated successfully")

		// Token is valid, proceed
		c.Next()
	}
}
