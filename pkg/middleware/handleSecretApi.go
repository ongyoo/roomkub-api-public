package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func ValidateSecret() gin.HandlerFunc {
	secretAPIKey := os.Getenv("SECRET_KEY")
	return func(c *gin.Context) {
		userSecretKey := c.GetHeader("secret-key")

		if secretAPIKey != userSecretKey {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message":           "Secret Not Found ",
				"secret-key-server": secretAPIKey,
				"secret-key-client": userSecretKey,
			})
			return
		}
	}
}
