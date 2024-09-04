package auth

import (
	"strings"

	"github.com/ekota-space/zero/pkgs/auth"
	"github.com/ekota-space/zero/pkgs/common"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func GetRefresh(ctx *gin.Context) {
	accessToken := strings.Split(ctx.GetHeader("Authorization"), "Bearer ")[1]

	if accessToken == "" {
		ctx.JSON(400, gin.H{"error": "Access token is required"})
		return
	}

	_, err := jwt.ParseWithClaims(accessToken, &auth.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(common.Env.JwtAccessTokenSecret), nil
	})

	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid access token"})
		return
	}

	// if jwtToken.Valid {
	// 	ctx.JSON(400, gin.H{"error": "Access token is still valid"})
	// 	return
	// }

	refreshToken := ctx.GetHeader("X-Refresh-Token")

	if refreshToken == "" {
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

	tokens.User = nil

	ctx.JSON(200, tokens)
}
