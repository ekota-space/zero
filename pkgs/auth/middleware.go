package auth

import (
	"github.com/ekota-space/zero/pkgs/response"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		accessToken := c.Cookies("acc_t")
		if accessToken == "" {
			return c.Status(401).JSON(response.Error("Access token is required"))
		}

		claims, _, err := VerifyAccessToken(accessToken)

		if err != nil {
			return c.Status(401).JSON(response.Error("Invalid access token"))
		}

		yes, userId, err := CheckUserExistsByEmail(claims.Email)

		if err != nil || !yes {
			if !yes {
				return c.Status(201).JSON(response.Error("User not found"))
			} else {
				return c.Status(201).JSON(response.Error("Invalid access token"))
			}
		}

		c.Locals("id", userId)
		c.Locals("username", claims.Username)
		c.Locals("email", claims.Email)
		return c.Next()
	}
}
