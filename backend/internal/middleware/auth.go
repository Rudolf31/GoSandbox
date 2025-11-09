package middleware

import (
	"interface_lesson/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewAuthMiddleware(auth services.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {

		username := c.GetHeader("X-Username")
		password := c.GetHeader("X-Password")

		if username == "" || password == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Missing credentials",
			})
			c.Abort()
			return
		}

		accountUsername, err := auth.Authenticate(username, password)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			c.Abort()
			return
		}

		c.Set("username", accountUsername)
		c.Next()
	}
}
