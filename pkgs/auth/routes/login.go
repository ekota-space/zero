package authRoutes

import (
	"github.com/ekota-space/zero/pkgs/auth"
	authDao "github.com/ekota-space/zero/pkgs/auth/dao"
	"github.com/ekota-space/zero/pkgs/common"
	"github.com/ekota-space/zero/pkgs/response"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// @Summary		Login user
// @Description	User login route
// @Tags			Auth
// @Accept			json
// @Produce		json
// @Param			body	body		authDao.LoginDao				true	"User login"
// @Success		200		{object}	authDao.AuthResponse			"User logged in"
// @Failure		400		{object}	response.ErrorResponse[string]	"Invalid email or password"
// @Failure		500		{object}	response.ErrorResponse[string]	"Internal server error"
// @Router			/auth/login [post]
func PostLogin(ctx *fiber.Ctx) error {
	body := authDao.LoginDao{}

	if err := ctx.BodyParser(&body); err != nil {
		return err
	}

	user, err := auth.GetUserByEmailUnsafely(body.Email)

	if err != nil {
		return ctx.Status(400).JSON(response.Error("Invalid email or password"))
	}

	if err := bcrypt.CompareHashAndPassword([]byte(*user.Password), []byte(body.Password)); err != nil {
		return ctx.Status(400).JSON(response.Error("Invalid email or password"))
	}

	tokens, err := auth.GenerateAuthTokens(user)

	if err != nil {
		return ctx.Status(500).JSON(response.Error(err.Error()))
	}

	auth.SetCookies(ctx, tokens)

	return ctx.JSON(
		authDao.AuthResponse{
			ExpirationDurationSeconds: int(common.AccessTokenDuration.Seconds()),
		},
	)
}
