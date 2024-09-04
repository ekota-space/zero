package auth

import (
	"strings"

	"github.com/ekota-space/zero/pkgs/auth"
	"github.com/ekota-space/zero/pkgs/common"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func GetRefresh(ctx *gin.Context) {
	accessToken, err := ctx.Cookie("acc_t")

	if err != nil || strings.TrimSpace(accessToken) == "" {
		ctx.JSON(400, gin.H{"error": "Access token is required"})
		return
	}

	_, jwtToken, err := auth.VerifyAccessToken(accessToken)

	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid access token"})
		return
	}

	if jwtToken.Valid {
		ctx.JSON(400, gin.H{"error": "Access token is still valid"})
		return
	}

	refreshToken, err := ctx.Cookie("ref_t")

	if err != nil || refreshToken == "" {
		ctx.JSON(400, gin.H{"error": "Refresh token is required"})
		return
	}

	refreshJwtToken, err := jwt.ParseWithClaims(refreshToken, &auth.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(common.Env.JwtRefreshTokenSecret), nil
	})

	if err != nil || !refreshJwtToken.Valid {
		ctx.JSON(400, gin.H{"error": "Invalid refresh token"})
		return
	}

	claims, ok := refreshJwtToken.Claims.(*auth.Claims)

	if !ok {
		ctx.JSON(400, gin.H{"error": "Invalid refresh token"})
	}

	user, err := auth.GetUserByEmail(claims.Email)

	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid refresh token"})
		return
	}

	tokens, err := auth.GenerateAuthTokens(&user)

	if err != nil {
		ctx.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	auth.SetCookies(ctx, tokens)

	ctx.Status(200)
}
