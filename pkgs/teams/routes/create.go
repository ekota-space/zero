package teamsRoutes

import (
	"github.com/ekota-space/zero/pkgs/common"
	"github.com/ekota-space/zero/pkgs/root/db/zero/public/model"
	"github.com/ekota-space/zero/pkgs/root/db/zero/public/table"
	"github.com/ekota-space/zero/pkgs/root/ql"
	teamsDao "github.com/ekota-space/zero/pkgs/teams/dao"
	"github.com/gin-gonic/gin"
)

func PostCreate(ctx *gin.Context) {
	body := teamsDao.CreateTeamInput{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	tx, err := ql.GetDB().Begin()

	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to start transaction"})
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
			ctx.JSON(400, gin.H{"error": "Slug already exists"})
			return
		}

		ctx.JSON(500, gin.H{"error": "Failed to create team"})
		return
	}

	tx.Commit()

	ctx.JSON(200, gin.H{"data": team})
}
