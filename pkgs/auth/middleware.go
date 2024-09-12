package auth

import (
	"github.com/gin-gonic/gin"
)

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

		yes, userId, err := CheckUserExistsByEmail(claims.Email)

		if err != nil || !yes {
			if !yes {
				c.JSON(401, gin.H{"error": "User not found"})
			} else {
				c.JSON(401, gin.H{"error": "Invalid access token"})
			}
			c.Abort()
			return
		}

		c.Set("id", userId)
		c.Set("username", claims.Username)
		c.Set("email", claims.Email)
		c.Next()
	}
}
