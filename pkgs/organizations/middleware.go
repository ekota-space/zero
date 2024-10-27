package organizations

import (
	"github.com/ekota-space/zero/pkgs/response"
	"github.com/gofiber/fiber/v2"
)

func AccessCheckMiddleware(accessRole string) fiber.Handler {

	return func(ctx *fiber.Ctx) error {
		userId := ctx.Locals("id").(string)
		orgSlug := ctx.Params("orgSlug")

		accessLevel, org, err := GetAccessLevel(userId, orgSlug)

		if err != nil {
			return ctx.Status(500).JSON(response.Error("internal server error"))

		}
		// Owner can access everything
		isOwner := accessLevel == OWNER && (accessRole == ADMIN || accessRole == MEMBER)
		// Admin can access everything except owner's
		isAdmin := accessLevel == ADMIN && accessRole == MEMBER

		ctx.Set("organizationId", org.ID.String())

		if accessLevel == accessRole || isOwner || isAdmin {
			return ctx.Next()

		}

		return ctx.Status(403).JSON(response.Error("forbidden"))
	}
}
