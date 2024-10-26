package authRoutes

import (
	"strings"

	"github.com/ekota-space/zero/pkgs/auth"
	authDao "github.com/ekota-space/zero/pkgs/auth/dao"
	"github.com/ekota-space/zero/pkgs/common"
	"github.com/ekota-space/zero/pkgs/response"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

//	@Summary		Refresh access token
//	@Description	Refresh access token route
//	@Tags			Auth
//	@Produce		json
//	@Success		200	{object}	authDao.AuthResponse			"Access token refreshed"
//	@Failure		400	{object}	response.ErrorResponse[string]	"Access token is required"
//	@Failure		400	{object}	response.ErrorResponse[string]	"Invalid access token"
//	@Failure		400	{object}	response.ErrorResponse[string]	"Access token is still valid"
//	@Failure		400	{object}	response.ErrorResponse[string]	"Refresh token is required"
//	@Failure		400	{object}	response.ErrorResponse[string]	"Invalid refresh token"
//	@Failure		500	{object}	response.ErrorResponse[string]	"Internal server error"
//	@Router			/auth/refresh [get]
func GetRefresh(ctx *gin.Context) {
	accessToken, err := ctx.Cookie("acc_t")

	if err != nil || strings.TrimSpace(accessToken) == "" {
		ctx.JSON(400, response.Error("Access token is required"))
		return
	}

	_, jwtToken, err := auth.VerifyAccessToken(accessToken)

	if err != nil {
		ctx.JSON(400, response.Error("Invalid access token"))
		return
	}

	if jwtToken.Valid {
		ctx.JSON(400, gin.H{
			"error":            "Access token is still valid",
			"expiresAtSeconds": jwtToken.Claims.(*auth.Claims).ExpiresAt.Unix(),
		})
		return
	}

	refreshToken, err := ctx.Cookie("ref_t")

	if err != nil || refreshToken == "" {
		ctx.JSON(400, response.Error("Refresh token is required"))
		return
	}

	refreshJwtToken, err := jwt.ParseWithClaims(refreshToken, &auth.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(common.Env.JwtRefreshTokenSecret), nil
	})

	if err != nil || !refreshJwtToken.Valid {
		ctx.JSON(400, response.Error("Invalid refresh token"))
		return
	}

	claims, ok := refreshJwtToken.Claims.(*auth.Claims)

	if !ok {
		ctx.JSON(400, response.Error("Invalid refresh token"))
	}

	user, err := auth.GetUserByEmail(claims.Email)

	if err != nil {
		ctx.JSON(400, response.Error("Invalid refresh token"))
		return
	}

	tokens, err := auth.GenerateAuthTokens(user)

	if err != nil {
		ctx.JSON(500, response.Error("Internal server error"))
		return
	}

	auth.SetCookies(ctx, tokens)

	ctx.JSON(200, authDao.AuthResponse{ExpirationDurationSeconds: int(common.AccessTokenDuration.Seconds())})
}
