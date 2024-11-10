package teamMembers

import (
	"github.com/ekota-space/zero/pkgs/response"
	"github.com/ekota-space/zero/pkgs/root/db/zero/public/model"
	"github.com/ekota-space/zero/pkgs/root/db/zero/public/table"
	"github.com/ekota-space/zero/pkgs/root/ql"
	jet "github.com/go-jet/jet/v2/postgres"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// @Summary		Get team members
// @Description	Get team members
// @Tags			Teams
// @Produce		json
// @Param			orgSlug	path	string	true	"Organization slug"
// @Param			teamSlug	path	string	true	"Team slug"
// @Success		200		{object}	response.SuccessDataResponse[[]model.TeamMembers]	"Team members"
// @Failure		404		{object}	response.ErrorResponse[string]			"Team members not found"
// @Failure		500		{object}	response.ErrorResponse[string]			"Failed to get team members"
// @Router		/organizations/{orgSlug}/teams/{teamSlug}/members [get]
func GetList(ctx *fiber.Ctx) error {
	orgId := ctx.Locals("organizationId").(string)
	teamSlug := ctx.Params("teamSlug")

	teamMembers := []model.TeamMembers{}

	stmt := table.TeamMembers.
		LEFT_JOIN(
			table.Users,
			table.Users.ID.EQ(table.TeamMembers.UserID),
		).
		LEFT_JOIN(
			table.Teams,
			table.Teams.ID.EQ(table.TeamMembers.TeamID),
		).
		SELECT(
			table.TeamMembers.AllColumns,
			table.Users.AllColumns,
		).WHERE(
		table.Teams.Slug.EQ(jet.String(teamSlug)).AND(
			table.Teams.OrganizationID.EQ(jet.UUID(uuid.MustParse(orgId))),
		),
	)

	err := stmt.Query(ql.GetDB(), &teamMembers)

	if err != nil {
		return ctx.Status(500).JSON(response.Error(err))
	}

	if len(teamMembers) == 0 {
		return ctx.Status(404).JSON(response.Error("Team members not found"))
	}

	return ctx.JSON(response.Success(teamMembers))
}
