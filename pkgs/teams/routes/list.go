package teamsRoutes

import (
	"github.com/ekota-space/zero/pkgs/root/db"
	"github.com/ekota-space/zero/pkgs/root/db/zero/public/model"
	"github.com/ekota-space/zero/pkgs/root/db/zero/public/table"
	"github.com/gin-gonic/gin"
	jet "github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"
)

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

	err := stmt.Query(db.DB, &teams)

	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to fetch teams"})
		return
	}

	ctx.JSON(200, gin.H{"data": teams})
}
