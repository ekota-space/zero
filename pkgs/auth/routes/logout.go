package authRoutes

import (
	"github.com/ekota-space/zero/pkgs/auth"
	"github.com/gofiber/fiber/v2"
)

// @Summary		Logout user
// @Description	User logout route
// @Tags			Auth
// @Produce		json
// @Success		200	"User logged out"
// @Router			/auth/logout [get]
func GetLogout(ctx *fiber.Ctx) error {
	auth.ClearCookies(ctx)

	return ctx.Status(200).JSON(nil)
}
