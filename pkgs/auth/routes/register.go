package authRoutes

import (
	"fmt"
	"strings"

	auth "github.com/ekota-space/zero/pkgs/auth"
	authDao "github.com/ekota-space/zero/pkgs/auth/dao"
	"github.com/ekota-space/zero/pkgs/common"
	"github.com/ekota-space/zero/pkgs/root/db/zero/public/model"
	"github.com/ekota-space/zero/pkgs/root/db/zero/public/table"
	"github.com/ekota-space/zero/pkgs/root/ql"

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

	user := model.Users{
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Username:  body.Username,
		Email:     body.Email,
		Password:  &passwordStr,
	}

	tx, err := ql.GetDB().Begin()

	if err != nil {
		fmt.Println(err)
		ctx.JSON(500, gin.H{"error": "Failed to start transaction"})
		return
	}

	stmt := table.Users.INSERT(
		table.Users.FirstName,
		table.Users.LastName,
		table.Users.Username,
		table.Users.Email,
		table.Users.Password,
	).
		MODEL(user).
		RETURNING(table.Users.AllColumns)

	newUser := model.Users{}

	err = stmt.Query(tx, &newUser)

	if err != nil {
		tx.Rollback()

		if strings.Contains(err.Error(), "email") {
			ctx.JSON(400, gin.H{"error": "Email already exists"})
		} else if strings.Contains(err.Error(), "username") {
			ctx.JSON(400, gin.H{"error": "Username already exists"})
		} else {
			ctx.JSON(500, gin.H{"error": err.Error()})
		}
		return
	}

	tx.Commit()

	// Generate Access Token and Refresh Token
	tokens, err := auth.GenerateAuthTokens(&user)

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	auth.SetCookies(ctx, tokens)

	ctx.JSON(201, gin.H{"expirationDurationSeconds": int(common.AccessTokenDuration.Seconds())})
}
