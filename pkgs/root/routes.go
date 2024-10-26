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
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()
	apiV1 := router.Group("/api/v1")

	router.Use(cors.New(
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
	apiV1.GET("/", root.GetRoot)

	{
		apiV1.POST("/auth/login", authRoutes.PostLogin)
		apiV1.POST("/auth/register", authRoutes.PostRegister)
		apiV1.GET("/auth/refresh", authRoutes.GetRefresh)
		apiV1.GET("/auth/logout", authRoutes.GetLogout)
	}

	{
		apiV1.Use(auth.AuthMiddleware())

		apiV1.GET("/user/me", userRoutes.GetUserInfo)

		{
			apiV1.GET("/organizations", organizationRoutes.GetList)
			apiV1.POST("/organizations", organizationRoutes.PostCreate)
		}

		{
			apiV1.POST(
				"/organizations/:orgSlug/teams",
				organizations.AccessCheckMiddleware(organizations.ADMIN),
				teamsRoutes.PostCreate,
			)
			apiV1.GET(
				"/organizations/:orgSlug/teams",
				organizations.AccessCheckMiddleware(organizations.MEMBER),
				teamsRoutes.GetList,
			)
		}
	}

	router.GET("/swagger/*any",
		func(ctx *gin.Context) {
			if ctx.Request.RequestURI == "/swagger/" {
				ctx.Redirect(302, "/swagger/index.html")
				return
			}
			ginSwagger.WrapHandler(swaggerFiles.Handler)(ctx)
		},
	)

	return router
}
