package root

import (
	auth "github.com/ekota-space/zero/pkgs/auth"
	authRoutes "github.com/ekota-space/zero/pkgs/auth/routes"
	"github.com/ekota-space/zero/pkgs/organizations"
	organizationRoutes "github.com/ekota-space/zero/pkgs/organizations/routes"
	root "github.com/ekota-space/zero/pkgs/root/routes"
	teamsRoutes "github.com/ekota-space/zero/pkgs/teams/routes"
	userRoutes "github.com/ekota-space/zero/pkgs/user/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
)

func SetupRoutes() *fiber.App {
	router := fiber.New()

	router.Use(logger.New())

	apiV1 := router.Group("/api/v1")

	apiV1.Use(cors.New(
		cors.Config{
			AllowOrigins:     "http://localhost:3000",
			AllowCredentials: true,
		},
	))
	apiV1.Get("/", root.GetRoot)

	{
		apiV1.Post("/auth/login", authRoutes.PostLogin)
		apiV1.Post("/auth/register", authRoutes.PostRegister)
		apiV1.Get("/auth/refresh", authRoutes.GetRefresh)
		apiV1.Get("/auth/logout", authRoutes.GetLogout)
	}

	{
		apiV1.Use(auth.AuthMiddleware())

		apiV1.Get("/user/me", userRoutes.GetUserInfo)

		{
			apiV1.Get("/organizations", organizationRoutes.GetList)
			apiV1.Post("/organizations", organizationRoutes.PostCreate)
		}

		{
			apiV1.Post(
				"/organizations/:orgSlug/teams",
				organizations.AccessCheckMiddleware(organizations.ADMIN),
				teamsRoutes.PostCreate,
			)
			apiV1.Get(
				"/organizations/:orgSlug/teams",
				organizations.AccessCheckMiddleware(organizations.MEMBER),
				teamsRoutes.GetList,
			)
		}
	}

	router.Get("/swagger/*", swagger.HandlerDefault)

	return router
}
