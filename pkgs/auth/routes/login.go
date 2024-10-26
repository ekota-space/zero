package authRoutes

import (
	"github.com/ekota-space/zero/pkgs/auth"
	authDao "github.com/ekota-space/zero/pkgs/auth/dao"
	"github.com/ekota-space/zero/pkgs/common"
	"github.com/ekota-space/zero/pkgs/response"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

//	@Summary		Login user
//	@Description	User login route
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			body	body		authDao.LoginDao				true	"User login"
//	@Success		200		{object}	authDao.AuthResponse			"User logged in"
//	@Failure		400		{object}	response.ErrorResponse[string]	"Invalid email or password"
//	@Failure		500		{object}	response.ErrorResponse[string]	"Internal server error"
//	@Router			/auth/login [post]
func PostLogin(ctx *gin.Context) {
	body := authDao.LoginDao{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(400, response.Error(err.Error()))
		return
	}

	user, err := auth.GetUserByEmailUnsafely(body.Email)

	if err != nil {
		ctx.JSON(400, response.Error("Invalid email or password"))
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(*user.Password), []byte(body.Password)); err != nil {
		ctx.JSON(400, response.Error("Invalid email or password"))
		return
	}

	tokens, err := auth.GenerateAuthTokens(user)

	if err != nil {
		ctx.JSON(500, response.Error(err.Error()))
		return
	}

	auth.SetCookies(ctx, tokens)

	ctx.JSON(
		200,
		authDao.AuthResponse{
			ExpirationDurationSeconds: int(common.AccessTokenDuration.Seconds()),
		},
	)
}
