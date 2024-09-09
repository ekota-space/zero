package authRoutes

import (
	"fmt"
	"strings"

	auth "github.com/ekota-space/zero/pkgs/auth"
	authDao "github.com/ekota-space/zero/pkgs/auth/dao"
	authModels "github.com/ekota-space/zero/pkgs/auth/models"
	"github.com/ekota-space/zero/pkgs/common"
	"github.com/ekota-space/zero/pkgs/root/db"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func PostRegister(ctx *gin.Context) {
	body := authDao.RegisterDao{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	password, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	passwordStr := string(password)

	user := authModels.Users{
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Username:  body.Username,
		Email:     body.Email,
		Password:  &passwordStr,
	}

	if err := db.DB.Create(&user).Error; err != nil {
		field := "Username"
		if strings.Contains(err.Error(), "email") {
			field = "Email"
		}
		ctx.JSON(400, gin.H{"error": fmt.Sprintf("%s already exists", field)})
		return
	}

	// Generate Access Token and Refresh Token
	tokens, err := auth.GenerateAuthTokens(&user)

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	auth.SetCookies(ctx, tokens)

	ctx.JSON(201, gin.H{"expirationDurationSeconds": int(common.AccessTokenDuration.Seconds())})
}
