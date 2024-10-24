package middleware

import (
	"edc-proxy/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the JWT token from the request header
		bearerTokenString := c.GetHeader("Authorization")

		// Log the request with the token (if present)
		utils.Log.WithFields(logrus.Fields{
			"token": bearerTokenString,
		}).Debug("Validating JWT")

		if bearerTokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token not provided"})
			c.Abort()
			return
		}

		authTokenParts := strings.Split(bearerTokenString, " ")
		if len(authTokenParts) != 2 || strings.ToLower(authTokenParts[0]) != "bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Not a valid Bearer Authorization token"})
			c.Abort()
			return
		}

		jwtString := authTokenParts[1]

		// Introspect the token
		token, err := utils.IntrospectJWT(jwtString)
		if err != nil {
			// Log the error
			utils.Log.WithFields(logrus.Fields{
				"error": err.Error(),
				"token": jwtString,
			}).Error("JWT validation failed")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Log successful validation
		utils.Log.WithFields(logrus.Fields{
			"token": jwtString,
		}).Debug("JWT token validated successfully")

		subject, err := utils.GetTokenClaim(token, "sub")
		if err != nil {
			utils.Log.WithFields(logrus.Fields{
				"error": err.Error(),
				"token": jwtString,
			}).Error("sub claim error")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
		organization, err := utils.GetTokenClaim(token, "organization")
		if err != nil {
			utils.Log.WithFields(logrus.Fields{
				"error": err.Error(),
				"token": jwtString,
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
