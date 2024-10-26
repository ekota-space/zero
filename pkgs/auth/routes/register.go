package authRoutes

import (
	"fmt"
	"strings"

	auth "github.com/ekota-space/zero/pkgs/auth"
	authDao "github.com/ekota-space/zero/pkgs/auth/dao"
	"github.com/ekota-space/zero/pkgs/common"
	"github.com/ekota-space/zero/pkgs/response"
	"github.com/ekota-space/zero/pkgs/root/db/zero/public/model"
	"github.com/ekota-space/zero/pkgs/root/db/zero/public/table"
	"github.com/ekota-space/zero/pkgs/root/ql"

	"github.com/gin-gonic/gin"
)

//	@Summary		Register user
//	@Description	User registration route
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			body	body		authDao.RegisterDao				true	"User registration"
//	@Success		201		{object}	authDao.AuthResponse			"User registered"
//	@Failure		400		{object}	response.ErrorResponse[string]	"Invalid request"
//	@Failure		500		{object}	response.ErrorResponse[string]	"Internal server error"
//	@Router			/auth/register [post]
func PostRegister(ctx *gin.Context) {
	body := authDao.RegisterDao{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(400, response.Error(err.Error()))
		return
	}

	passwordStr, err := auth.HashPassword(body.Password)

	if err != nil {
		ctx.JSON(500, response.Error(err.Error()))
		return
	}

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
		ctx.JSON(500, response.Error("Failed to start transaction"))
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
			ctx.JSON(400, response.Error("Email already exists"))
		} else if strings.Contains(err.Error(), "username") {
			ctx.JSON(400, response.Error("Username already exists"))
		} else {
			ctx.JSON(500, response.Error(err.Error()))
		}
		return
	}

	tx.Commit()

	// Generate Access Token and Refresh Token
	tokens, err := auth.GenerateAuthTokens(&user)

	if err != nil {
		ctx.JSON(500, response.Error(err.Error()))
		return
	}

	auth.SetCookies(ctx, tokens)

	ctx.JSON(201, authDao.AuthResponse{ExpirationDurationSeconds: int(common.AccessTokenDuration.Seconds())})
}
