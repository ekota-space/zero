package authRoutes

import (
	"github.com/ekota-space/zero/pkgs/auth"
	authDao "github.com/ekota-space/zero/pkgs/auth/dao"
	"github.com/ekota-space/zero/pkgs/common"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func PostLogin(ctx *gin.Context) {
	body := authDao.LoginDao{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, err := auth.GetUserByEmailUnsafely(body.Email)

	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid email or password"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(*user.Password), []byte(body.Password)); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid email or password"})
		return
	}

	tokens, err := auth.GenerateAuthTokens(&user)

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	auth.SetCookies(ctx, tokens)

	ctx.JSON(200, gin.H{"expirationDurationSeconds": int(common.AccessTokenDuration.Seconds())})
}
