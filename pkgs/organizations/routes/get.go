package organizationRoutes

import (
	"github.com/ekota-space/zero/pkgs/response"
	"github.com/ekota-space/zero/pkgs/root/db/zero/public/model"
	"github.com/ekota-space/zero/pkgs/root/db/zero/public/table"
	"github.com/ekota-space/zero/pkgs/root/ql"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	jet "github.com/go-jet/jet/v2/postgres"
)

// @Summary		Get organization
// @Description	Get organization by slug
// @Tags			Organizations
// @Accept			json
// @Produce		json
// @Param			orgId	path		string												true	"Organization Slug"
// @Success		200		{object}	response.SuccessDataResponse[model.Organizations]	"Successful response"
// @Failure		404		{object}	response.ErrorResponse[string]						"Organization not found"
// @Router			/organizations/{orgId} [get]
func GetOrganization(ctx *fiber.Ctx) error {
	slug := ctx.Params("orgId")

	organization := model.Organizations{}

	stmt := table.Organizations.SELECT(table.Organizations.AllColumns).WHERE(
		table.Organizations.Slug.EQ(
			jet.String(slug),
		),
	)

	err := stmt.Query(ql.GetDB(), &organization)

	if err != nil || organization.ID == uuid.Nil {
		return ctx.Status(404).JSON(response.Error("Organization not found"))
	}

	return ctx.Status(200).JSON(response.Success(organization))
}
