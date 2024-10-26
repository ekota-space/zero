package teamsRoutes

import (
	"github.com/ekota-space/zero/pkgs/common"
	"github.com/ekota-space/zero/pkgs/response"
	"github.com/ekota-space/zero/pkgs/root/db/zero/public/model"
	"github.com/ekota-space/zero/pkgs/root/db/zero/public/table"
	"github.com/ekota-space/zero/pkgs/root/ql"
	teamsDao "github.com/ekota-space/zero/pkgs/teams/dao"
	"github.com/gin-gonic/gin"
)

// @Summary		Create a team
// @Description	Create a team
// @Tags			Teams
// @Accept			json
// @Produce		json
// @Param			orgSlug	path	string	true	"Organization slug"
// @Param			body	body		teamsDao.CreateTeamInput				true	"Team data"
// @Success		200		{object}	response.SuccessDataResponse[model.Teams]	"Team created"
// @Failure		400		{object}	response.ErrorResponse[string]			"Slug already exists"
// @Failure		500		{object}	response.ErrorResponse[string]			"Failed to create team"
// @Router		/organizations/{orgSlug}/teams [post]
func PostCreate(ctx *gin.Context) {
	body := teamsDao.CreateTeamInput{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(400, response.Error(err.Error()))
		return
	}

	tx, err := ql.GetDB().Begin()

	if err != nil {
		ctx.JSON(500, response.Error("Failed to start transaction"))
		return
	}

	stmt := table.Teams.INSERT(
		table.Teams.Name,
		table.Teams.Description,
		table.Teams.Slug,
	).
		MODEL(body).
		RETURNING(table.Teams.AllColumns)

	team := model.Teams{}

	err = stmt.Query(tx, &team)

	if err != nil {
		tx.Rollback()

		if common.IsDuplicateKeyError(err) {
			ctx.JSON(400, response.Error("Slug already exists"))
			return
		}

		ctx.JSON(500, response.Error("Failed to create team"))
		return
	}

	tx.Commit()

	ctx.JSON(200, response.Success(team))
}
