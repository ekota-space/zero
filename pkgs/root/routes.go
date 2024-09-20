package root

import (
	"time"

	auth "github.com/ekota-space/zero/pkgs/auth"
	authRoutes "github.com/ekota-space/zero/pkgs/auth/routes"
	"github.com/ekota-space/zero/pkgs/organizations"
	organizationRoutes "github.com/ekota-space/zero/pkgs/organizations/routes"
	root "github.com/ekota-space/zero/pkgs/root/routes"
	teamsRoutes "github.com/ekota-space/zero/pkgs/teams/routes"
	userRoutes "github.com/ekota-space/zero/pkgs/user/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
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

	orgs := protected.Group("/organizations")
	{
		orgs.GET("/", organizations.AccessCheckMiddleware(organizations.MEMBER), organizationRoutes.GetList)
		orgs.POST("/", organizationRoutes.PostCreate)
	}

	teams := orgs.Group("/:orgSlug/teams")
	{
		teams.POST("/", organizations.AccessCheckMiddleware(organizations.ADMIN), teamsRoutes.PostCreate)
		teams.GET("/", organizations.AccessCheckMiddleware(organizations.MEMBER), teamsRoutes.GetList)
	}

	return r
}
