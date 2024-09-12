package organizationRoutes

import (
	"github.com/ekota-space/zero/pkgs/root/db"
	"github.com/ekota-space/zero/pkgs/root/db/zero/public/model"
	"github.com/ekota-space/zero/pkgs/root/db/zero/public/table"
	"github.com/gin-gonic/gin"
	jet "github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"
)

func GetList(ctx *gin.Context) {
	userId := ctx.GetString("id")

	stmt := table.Organizations.
		SELECT(table.Organizations.AllColumns).
		WHERE(table.Organizations.OwnerID.EQ(jet.UUID(uuid.MustParse(userId))))

	organizations := []model.Organizations{}

	err := stmt.Query(db.DB, &organizations)

	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"data": organizations})
}
