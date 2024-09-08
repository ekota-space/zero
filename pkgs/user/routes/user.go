package userRoutes

import (
	"github.com/ekota-space/zero/pkgs/auth"
	"github.com/gin-gonic/gin"
)

func GetUserInfo(ctx *gin.Context) {
	email := ctx.GetString("email")

	user, err := auth.GetUserByEmail(email)

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, user)
}
