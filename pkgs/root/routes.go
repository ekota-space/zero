package root

import (
	"fmt"
	"time"

	auth "github.com/ekota-space/zero/pkgs/auth"
	authRoutes "github.com/ekota-space/zero/pkgs/auth/routes"
	"github.com/ekota-space/zero/pkgs/common"
	organizationRoutes "github.com/ekota-space/zero/pkgs/organizations/routes"
	root "github.com/ekota-space/zero/pkgs/root/routes"
	userRoutes "github.com/ekota-space/zero/pkgs/user/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() {
	r := gin.Default()

	corsConfig := cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization", "Access-Control-Allow-Headers"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	r.Use(cors.New(corsConfig))

	r.GET("/", root.GetRoot)

	r.POST("/auth/login", authRoutes.PostLogin)
	r.POST("/auth/register", authRoutes.PostRegister)
	r.GET("/auth/refresh", authRoutes.GetRefresh)
	r.GET("/auth/logout", authRoutes.GetLogout)

	protected := r.Group("/")

	protected.Use(auth.AuthMiddleware())

	protected.GET("/user/me", userRoutes.GetUserInfo)

	protected.GET("/organizations", organizationRoutes.GetList)
	protected.POST("/organizations", organizationRoutes.PostCreate)

	protected.GET("/organizations/:id", organizationRoutes.GetOrganization)

	r.Run(fmt.Sprintf("localhost:%d", common.Env.Port))
}
