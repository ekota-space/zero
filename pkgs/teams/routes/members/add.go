package teamMembers

import (
	"github.com/ekota-space/zero/pkgs/response"
	teamsDao "github.com/ekota-space/zero/pkgs/teams/dao"
	"github.com/gofiber/fiber/v2"
)

// @Summary		Add a member to a team
// @Description	Add a member to a team
// @Tags			Teams
// @Accept		json
// @Produce		json
// @Param			orgSlug	path	string	true	"Organization slug"
// @Param			teamSlug	path	string	true	"Team slug"
// @Param			body	body		teamsDao.AddMemberInput				true	"Team member id"
// @Success		201		{object}	response.SuccessDataResponse[model.TeamMember]	"Team member added"
// @Failure		400		{object}	response.ErrorResponse[string]			"Team member already exists"
// @Failure		400		{object}	response.ErrorResponse[string]			"Failed to parse body"
// @Failure		500		{object}	response.ErrorResponse[string]			"Failed to add member to the team"
// @Router		/organizations/{orgSlug}/teams/{teamSlug}/members [post]
func PostAdd(ctx *fiber.Ctx) error {
	orgId := ctx.Locals("orgId")
	teamSlug := ctx.Params("teamSlug")

	body := teamsDao.AddMemberInput{}

	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(400).JSON(response.Error(err.Error()))
	}
}
