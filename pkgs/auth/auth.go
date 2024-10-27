package auth

import (
	"time"

	authDao "github.com/ekota-space/zero/pkgs/auth/dao"
	"github.com/ekota-space/zero/pkgs/common"
	"github.com/ekota-space/zero/pkgs/root/db/zero/public/model"
	"github.com/ekota-space/zero/pkgs/root/db/zero/public/table"
	"github.com/ekota-space/zero/pkgs/root/ql"
	jet "github.com/go-jet/jet/v2/postgres"
	"github.com/gofiber/fiber/v2"
	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

func GetUserByEmailUnsafely(email string) (*model.Users, error) {
	user := model.Users{}

	stmt := table.Users.SELECT(table.Users.AllColumns).WHERE(table.Users.Email.EQ(jet.String(email)))

	if err := stmt.Query(ql.GetDB(), &user); err != nil {
		return nil, err
	}
	return &user, nil
}

func CheckUserExistsByEmail(email string) (bool, string, error) {
	stmt := table.Users.SELECT(table.Users.ID).WHERE(table.Users.Email.EQ(jet.String(email)))

	user := model.Users{}
	if err := stmt.Query(ql.GetDB(), &user); err != nil {
		return false, "", err
	}

	return user.ID != uuid.Nil, user.ID.String(), nil
}

func GetUserByEmail(email string) (*model.Users, error) {
	user := model.Users{}

	stmt := table.Users.
		SELECT(table.Users.AllColumns.Except(table.Users.Password)).
		WHERE(table.Users.Email.EQ(jet.String(email)))

	if err := stmt.Query(ql.GetDB(), &user); err != nil {
		return nil, err
	}
	return &user, nil
}

type Claims struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateAuthTokens(user *model.Users) (authDao.AuthTokenResponseDao, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256,
		Claims{
			Username: user.Username,
			Email:    user.Email,
			RegisteredClaims: jwt.RegisteredClaims{
				Issuer:    "zero",
				Subject:   user.Username,
				Audience:  jwt.ClaimStrings{"zero"},
				ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(common.AccessTokenDuration)), // 1 hour
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
				ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(common.RefreshTokenDuration)), // 6 months
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

func SetCookies(ctx *fiber.Ctx, tokens authDao.AuthTokenResponseDao) {
	ctx.Cookie(
		&fiber.Cookie{
			Name:     "acc_t",
			Value:    tokens.AccessToken,
			MaxAge:   int(time.Hour * 24 * 30), /*30 days*/
			Path:     "/",
			Domain:   common.Env.ClientOrigin,
			Secure:   false,
			HTTPOnly: true,
		},
	)

	ctx.Cookie(
		&fiber.Cookie{
			Name:     "ref_t",
			Value:    tokens.RefreshToken,
			MaxAge:   int(time.Hour * 24 * 30), /*30 days*/
			Path:     "/",
			Domain:   common.Env.ClientOrigin,
			Secure:   false,
			HTTPOnly: true,
		},
	)
}

func ClearCookies(ctx *fiber.Ctx) {
	ctx.Cookie(
		&fiber.Cookie{
			Name:     "acc_t",
			Value:    "",
			MaxAge:   -1,
			Domain:   common.Env.ClientOrigin,
			Path:     "/",
			Secure:   false,
			HTTPOnly: true,
		},
	)
	ctx.Cookie(&fiber.Cookie{
		Name:     "ref_t",
		Value:    "",
		MaxAge:   -1,
		Domain:   common.Env.ClientOrigin,
		Path:     "/",
		Secure:   false,
		HTTPOnly: true,
	})
}
