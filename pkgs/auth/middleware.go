package auth

import (
	"github.com/ekota-space/zero/pkgs/response"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken, err := c.Cookie("acc_t")
		if err != nil || accessToken == "" {
			c.JSON(401, response.Error("Missing access token"))
			c.Abort()
			return
		}

		claims, _, err := VerifyAccessToken(accessToken)

		if err != nil {
			c.JSON(401, response.Error("Invalid access token"))
			c.Abort()
			return
		}

		yes, userId, err := CheckUserExistsByEmail(claims.Email)

		if err != nil || !yes {
			if !yes {
				c.JSON(401, response.Error("User not found"))
			} else {
				c.JSON(401, response.Error("Invalid access token"))
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
