package root

import (
	"net/http"

	auth "github.com/ekota-space/zero/pkgs/auth"
	authRoutes "github.com/ekota-space/zero/pkgs/auth/routes"
	"github.com/ekota-space/zero/pkgs/organizations"
	organizationRoutes "github.com/ekota-space/zero/pkgs/organizations/routes"
	root "github.com/ekota-space/zero/pkgs/root/routes"
	teamsRoutes "github.com/ekota-space/zero/pkgs/teams/routes"
	userRoutes "github.com/ekota-space/zero/pkgs/user/routes"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(
		cors.Options{
			AllowedOrigins: []string{"http://localhost:3000"},
			AllowedMethods: []string{
				http.MethodOptions,
				http.MethodHead,
				http.MethodGet,
				http.MethodPost,
				http.MethodPut,
				http.MethodPatch,
				http.MethodDelete,
			},
			AllowedHeaders:   []string{"*"},
			AllowCredentials: true,
		},
	))
	r.GET("/", root.GetRoot)

	{
		r.POST("/auth/login", authRoutes.PostLogin)
		r.POST("/auth/register", authRoutes.PostRegister)
		r.GET("/auth/refresh", authRoutes.GetRefresh)
		r.GET("/auth/logout", authRoutes.GetLogout)
	}

	{
		r.Use(auth.AuthMiddleware())

		r.GET("/user/me", userRoutes.GetUserInfo)

		{
			r.GET("/organizations", organizationRoutes.GetList)
			r.POST("/organizations", organizationRoutes.PostCreate)
		}

		{
			r.POST(
				"/organizations/:orgSlug/teams",
				organizations.AccessCheckMiddleware(organizations.ADMIN),
				teamsRoutes.PostCreate,
			)
			r.GET(
				"/organizations/:orgSlug/teams",
				organizations.AccessCheckMiddleware(organizations.MEMBER),
				teamsRoutes.GetList,
			)
		}
	}

	return r
}
