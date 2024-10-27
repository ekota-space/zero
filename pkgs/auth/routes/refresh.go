package authRoutes

import (
	"strings"

	"github.com/ekota-space/zero/pkgs/auth"
	authDao "github.com/ekota-space/zero/pkgs/auth/dao"
	"github.com/ekota-space/zero/pkgs/common"
	"github.com/ekota-space/zero/pkgs/response"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// @Summary		Refresh access token
// @Description	Refresh access token route
// @Tags			Auth
// @Produce		json
// @Success		200	{object}	authDao.AuthResponse			"Access token refreshed"
// @Failure		400	{object}	response.ErrorResponse[string]	"Access token is required"
// @Failure		400	{object}	response.ErrorResponse[string]	"Invalid access token"
// @Failure		400	{object}	response.ErrorResponse[string]	"Access token is still valid"
// @Failure		400	{object}	response.ErrorResponse[string]	"Refresh token is required"
// @Failure		400	{object}	response.ErrorResponse[string]	"Invalid refresh token"
// @Failure		500	{object}	response.ErrorResponse[string]	"Internal server error"
// @Router			/auth/refresh [get]
func GetRefresh(ctx *fiber.Ctx) error {
	accessToken := ctx.Cookies("acc_t")

	if strings.TrimSpace(accessToken) == "" {
		return ctx.Status(400).JSON(response.Error("Access token is required"))

	}

	_, jwtToken, err := auth.VerifyAccessToken(accessToken)

	if err != nil {
		return ctx.Status(400).JSON(response.Error("Invalid access token"))

	}

	if jwtToken.Valid {
		return ctx.Status(400).JSON(fiber.Map{
			"error":            "Access token is still valid",
			"expiresAtSeconds": jwtToken.Claims.(*auth.Claims).ExpiresAt.Unix(),
		})

	}

	refreshToken := ctx.Cookies("ref_t")

	if err != nil || refreshToken == "" {
		return ctx.Status(400).JSON(response.Error("Refresh token is required"))

	}

	refreshJwtToken, err := jwt.ParseWithClaims(refreshToken, &auth.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(common.Env.JwtRefreshTokenSecret), nil
	})

	if err != nil || !refreshJwtToken.Valid {
		return ctx.Status(400).JSON(response.Error("Invalid refresh token"))

	}

	claims, ok := refreshJwtToken.Claims.(*auth.Claims)

	if !ok {
		return ctx.Status(400).JSON(response.Error("Invalid refresh token"))
	}

	user, err := auth.GetUserByEmail(claims.Email)

	if err != nil {
		return ctx.Status(400).JSON(response.Error("Invalid refresh token"))

	}

	tokens, err := auth.GenerateAuthTokens(user)

	if err != nil {
		return ctx.Status(500).JSON(response.Error("Internal server error"))

	}

	auth.SetCookies(ctx, tokens)

	return ctx.Status(200).JSON(
		authDao.AuthResponse{
			ExpirationDurationSeconds: int(common.AccessTokenDuration.Seconds()),
		},
	)
}
