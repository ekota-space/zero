package auth

import "github.com/gin-gonic/gin"

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken, err := c.Cookie("acc_t")
		if err != nil || accessToken == "" {
			c.JSON(401, gin.H{"error": "Missing access token"})
			c.Abort()
			return
		}

		claims, _, err := VerifyAccessToken(accessToken)
		if err != nil {
			c.JSON(401, gin.H{"error": "Invalid access token"})
			c.Abort()
			return
		}

		c.Set("username", claims.Username)
		c.Set("email", claims.Email)
		c.Next()
	}
}