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
	"github.com/gofiber/fiber/v2"
)

// @Summary		Register user
// @Description	User registration route
// @Tags			Auth
// @Accept			json
// @Produce		json
// @Param			body	body		authDao.RegisterDao				true	"User registration"
// @Success		201		{object}	authDao.AuthResponse			"User registered"
// @Failure		400		{object}	response.ErrorResponse[string]	"Invalid request"
// @Failure		500		{object}	response.ErrorResponse[string]	"Internal server error"
// @Router			/auth/register [post]
func PostRegister(ctx *fiber.Ctx) error {
	body := authDao.RegisterDao{}

	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(400).JSON(response.Error(err.Error()))

	}

	passwordStr, err := auth.HashPassword(body.Password)

	if err != nil {
		return ctx.Status(500).JSON(response.Error(err.Error()))

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
		return ctx.Status(500).JSON(response.Error("Failed to start transaction"))

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
			return ctx.Status(400).JSON(response.Error("Email already exists"))
		} else if strings.Contains(err.Error(), "username") {
			return ctx.Status(400).JSON(response.Error("Username already exists"))
		} else {
			return ctx.Status(500).JSON(response.Error(err.Error()))
		}

	}

	tx.Commit()

	// Generate Access Token and Refresh Token
	tokens, err := auth.GenerateAuthTokens(&user)

	if err != nil {
		return ctx.Status(500).JSON(response.Error(err.Error()))
	}

	auth.SetCookies(ctx, tokens)

	return ctx.Status(201).JSON(
		authDao.AuthResponse{
			ExpirationDurationSeconds: int(common.AccessTokenDuration.Seconds()),
		},
	)
}
