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
		}).Debug("Validating JWT")

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token not provided"})
			c.Abort()
			return
		}

		// Introspect the token
		token, err := utils.IntrospectJWT(tokenString)
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
		}).Debug("JWT token validated successfully")

		subject, err := utils.GetTokenClaim(token, "sub")
		if err != nil {
			utils.Log.WithFields(logrus.Fields{
				"error": err.Error(),
				"token": tokenString,
			}).Error("sub claim error")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
		organization, err := utils.GetTokenClaim(token, "organization")
		if err != nil {
			utils.Log.WithFields(logrus.Fields{
				"error": err.Error(),
				"token": tokenString,
			}).Error("organization claim error")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
		c.Set("currentUserSub", subject)
		c.Set("currentUserOrg", organization)

		c.Next()
	}
}
