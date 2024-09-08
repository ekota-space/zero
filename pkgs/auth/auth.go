package auth

import (
	"time"

	authDao "github.com/ekota-space/zero/pkgs/auth/dao"
	auth "github.com/ekota-space/zero/pkgs/auth/models"
	"github.com/ekota-space/zero/pkgs/common"
	"github.com/ekota-space/zero/pkgs/root/db"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func GetUserByEmailUnsafely(email string) (auth.Users, error) {
	user := auth.Users{}
	if err := db.DB.Where("email = ?", email).First(&user); err.Error != nil {
		return auth.Users{}, err.Error
	}
	return user, nil
}

func GetUserByEmail(email string) (auth.Users, error) {
	user, err := GetUserByEmailUnsafely(email)

	if err != nil {
		return auth.Users{}, err
	}

	user.Password = nil
	return user, nil
}

type Claims struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateAuthTokens(user *auth.Users) (authDao.AuthTokenResponseDao, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256,
		Claims{
			Username: user.Username,
			Email:    user.Email,
			RegisteredClaims: jwt.RegisteredClaims{
				Issuer:    "zero",
				Subject:   user.Username,
				Audience:  jwt.ClaimStrings{"zero"},
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(common.AccessTokenDuration)), // 1 hour
			},
		},
	)

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256,
		Claims{
			Username: user.Username,
			Email:    user.Email,
			RegisteredClaims: jwt.RegisteredClaims{
				Issuer:    "zero",
				Subject:   user.Username,
				Audience:  jwt.ClaimStrings{"zero"},
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(common.RefreshTokenDuration)), // 6 months
			},
		},
	)

	accessTokenString, err := accessToken.SignedString([]byte(common.Env.JwtAccessTokenSecret))

	if err != nil {
		return authDao.AuthTokenResponseDao{}, err
	}

	refreshTokenString, err := refreshToken.SignedString([]byte(common.Env.JwtRefreshTokenSecret))

	if err != nil {
		return authDao.AuthTokenResponseDao{}, err
	}

	return authDao.AuthTokenResponseDao{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}, nil
}

func VerifyAccessToken(token string) (*Claims, *jwt.Token, error) {
	claims := &Claims{}
	accessToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(common.Env.JwtAccessTokenSecret), nil
	})

	if err != nil {
		return nil, nil, err
	}

	return claims, accessToken, nil

}

func SetCookies(ctx *gin.Context, tokens authDao.AuthTokenResponseDao) {
	ctx.SetCookie("acc_t", tokens.AccessToken, int(common.AccessTokenDuration), "/", common.Env.ClientOrigin, false, true)
	ctx.SetCookie("ref_t", tokens.RefreshToken, int(common.RefreshTokenDuration), "/", common.Env.ClientOrigin, false, true)
}

func ClearCookies(ctx *gin.Context) {
	ctx.SetCookie("acc_t", "", -1, "/", common.Env.ClientOrigin, false, true)
	ctx.SetCookie("ref_t", "", -1, "/", common.Env.ClientOrigin, false, true)
}
