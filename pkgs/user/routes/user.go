package userRoutes

import (
	"github.com/ekota-space/zero/pkgs/auth"
	"github.com/ekota-space/zero/pkgs/response"
	"github.com/gofiber/fiber/v2"
)

// @Summary		Get user info
// @Description	Get user info
// @Tags			User
// @Accept			json
// @Produce		json
// @Success		200	{object}	response.SuccessDataResponse[model.Users]	"User info"
// @Failure		500	{object}	response.ErrorResponse[string]	"Failed to fetch user info"
// @Router		/user/me [get]
func GetUserInfo(ctx *fiber.Ctx) error {
	email := ctx.Locals("email").(string)

	user, err := auth.GetUserByEmail(email)

	if err != nil {
		return ctx.Status(500).JSON(response.Error(err.Error()))
	}

	return ctx.Status(200).JSON(response.Success(user))
}
