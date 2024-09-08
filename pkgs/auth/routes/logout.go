package authRoutes

import (
	"github.com/ekota-space/zero/pkgs/auth"
	"github.com/gin-gonic/gin"
)

func GetLogout(ctx *gin.Context) {
	auth.ClearCookies(ctx)

	ctx.Status(200)
}
