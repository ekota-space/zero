package root

import (
	"fmt"

	auth "github.com/ekota-space/zero/pkgs/auth"
	authRoutes "github.com/ekota-space/zero/pkgs/auth/routes"
	"github.com/ekota-space/zero/pkgs/common"
	root "github.com/ekota-space/zero/pkgs/root/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() {
	r := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"} // TODO: Change this to a specific domain
	r.Use(cors.New(corsConfig))

	r.GET("/", root.GetRoot)

	r.POST("/auth/login", authRoutes.PostLogin)
	r.POST("/auth/register", authRoutes.PostRegister)
	r.GET("/auth/refresh", authRoutes.GetRefresh)
	r.GET("/auth/logout", authRoutes.GetLogout)

	protected := r.Group("/")

	protected.Use(auth.AuthMiddleware())

	r.Run(fmt.Sprintf("localhost:%d", common.Env.Port))
}
