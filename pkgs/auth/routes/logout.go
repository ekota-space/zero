package authRoutes

import (
	"github.com/ekota-space/zero/pkgs/auth"
	"github.com/gin-gonic/gin"
)

// @Summary		Logout user
// @Description	User logout route
// @Tags			Auth
// @Produce		json
// @Success		200	"User logged out"
// @Router			/auth/logout [get]
func GetLogout(ctx *gin.Context) {
	auth.ClearCookies(ctx)

	ctx.Status(200)
}
