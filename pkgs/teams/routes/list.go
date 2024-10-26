package teamsRoutes

import (
	"github.com/ekota-space/zero/pkgs/response"
	"github.com/ekota-space/zero/pkgs/root/db/zero/public/model"
	"github.com/ekota-space/zero/pkgs/root/db/zero/public/table"
	"github.com/ekota-space/zero/pkgs/root/ql"
	"github.com/gin-gonic/gin"
	jet "github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"
)

// @Summary		Get list of teams
// @Description	Get list of teams
// @Tags			Teams
// @Accept			json
// @Produce		json
// @Param			orgSlug	path	string	true	"Organization slug"
// @Success		200	{object}	response.SuccessDataResponse[[]model.Teams]	"List of teams"
// @Failure		500	{object}	response.ErrorResponse[string]			"Failed to fetch teams"
// @Router		/organizations/{orgSlug}/teams [get]
func GetList(ctx *gin.Context) {
	orgId := ctx.GetString("organizationId")

	teams := []model.Teams{}

	stmt := table.Teams.
		SELECT(
			table.Teams.AllColumns,
		).
		WHERE(
			table.Teams.OrganizationID.EQ(
				jet.UUID(uuid.MustParse(orgId)),
			),
		)

	err := stmt.Query(ql.GetDB(), &teams)

	if err != nil {
		ctx.JSON(500, response.Error("Failed to fetch teams"))
		return
	}

	ctx.JSON(200, response.Success(teams))
}
