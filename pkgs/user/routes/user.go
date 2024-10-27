package userRoutes

import (
	"github.com/ekota-space/zero/pkgs/auth"
	"github.com/ekota-space/zero/pkgs/response"
	"github.com/gin-gonic/gin"
)

// @Summary		Get user info
// @Description	Get user info
// @Tags			User
// @Accept			json
// @Produce		json
// @Success		200	{object}	response.SuccessDataResponse[model.Users]	"User info"
// @Failure		500	{object}	response.ErrorResponse[string]	"Failed to fetch user info"
// @Router		/user/me [get]
func GetUserInfo(ctx *gin.Context) {
	email := ctx.GetString("email")

	user, err := auth.GetUserByEmail(email)

	if err != nil {
		ctx.JSON(500, response.Error(err.Error()))
		return
	}

	ctx.JSON(200, response.Success(user))
}
