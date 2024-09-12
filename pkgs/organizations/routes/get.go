package organizationRoutes

import (
	"github.com/ekota-space/zero/pkgs/root/db"
	"github.com/ekota-space/zero/pkgs/root/db/zero/public/model"
	"github.com/ekota-space/zero/pkgs/root/db/zero/public/table"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	jet "github.com/go-jet/jet/v2/postgres"
)

func GetOrganization(ctx *gin.Context) {
	orgId := ctx.Param("id")

	organization := model.Organizations{}

	stmt := table.Organizations.SELECT(table.Organizations.AllColumns).WHERE(
		table.Organizations.ID.EQ(
			jet.UUID(uuid.MustParse(orgId)),
		),
	)

	err := stmt.Query(db.DB, &organization)

	if err != nil || organization.ID == uuid.Nil {
		ctx.JSON(404, gin.H{"error": "Organization not found"})
		return
	}

	ctx.JSON(200, gin.H{"data": organization})
}
