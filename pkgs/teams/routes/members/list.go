package teamMembers

import "github.com/gofiber/fiber/v2"

func GetList(ctx *fiber.Ctx) error {
	orgId := ctx.Locals("organizationId").(string)
	teamSlug := ctx.Params("teamSlug")

}
