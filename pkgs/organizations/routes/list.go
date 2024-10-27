package organizationRoutes

import (
	"github.com/ekota-space/zero/pkgs/response"
	"github.com/ekota-space/zero/pkgs/root/db/zero/public/model"
	"github.com/ekota-space/zero/pkgs/root/db/zero/public/table"
	"github.com/ekota-space/zero/pkgs/root/ql"
	jet "github.com/go-jet/jet/v2/postgres"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// @Summary		List organizations
// @Description	List organizations route
// @Tags			Organizations
// @Accept			json
// @Produce		json
// @Success		200	{object}	response.SuccessDataResponse[[]model.Organizations]	"List of organizations"
// @Failure		400	{object}	response.ErrorResponse[string]						"Invalid request"
// @Router			/organizations [get]
func GetList(ctx *fiber.Ctx) error {
	userId := ctx.Locals("id").(string)

	stmt := table.Organizations.
		LEFT_JOIN(
			table.OrganizationMembers,
			table.OrganizationMembers.OrganizationID.EQ(table.Organizations.ID).
				AND(table.OrganizationMembers.ID.EQ(jet.UUID(uuid.MustParse(userId)))),
		).
		SELECT(table.Organizations.AllColumns).
		WHERE(
			table.Organizations.OwnerID.EQ(jet.UUID(uuid.MustParse(userId))).
				OR(table.OrganizationMembers.ID.EQ(jet.UUID(uuid.MustParse(userId)))),
		)

	organizations := []model.Organizations{}

	err := stmt.Query(ql.GetDB(), &organizations)

	if err != nil {
		return ctx.Status(400).JSON(response.Error(err.Error()))
	}

	return ctx.Status(200).JSON(response.Success(organizations))
}
