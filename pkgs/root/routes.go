package root

import (
	"fmt"

	auth "github.com/ekota-space/zero/pkgs/auth/routes"
	"github.com/ekota-space/zero/pkgs/common"
	root "github.com/ekota-space/zero/pkgs/root/routes"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() {
	r := gin.Default()

	r.GET("/", root.GetRoot)

	r.POST("/auth/login", auth.PostLogin)
	r.POST("/auth/register", auth.PostRegister)
	r.GET("/auth/refresh", auth.GetRefresh)

	r.Run(fmt.Sprintf("localhost:%d", common.Env.Port))
}
