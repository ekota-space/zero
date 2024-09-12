package organizationRoutes

import (
	"strconv"

	organizationDao "github.com/ekota-space/zero/pkgs/organizations/dao"
	"github.com/ekota-space/zero/pkgs/root/db"
	"github.com/ekota-space/zero/pkgs/root/db/zero/public/table"
	"github.com/gin-gonic/gin"
	jet "github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"
)

func GetList(ctx *gin.Context) {
	offset, ok := ctx.GetQuery("offset")
	if !ok {
		offset = "0"
	}

	limit, ok := ctx.GetQuery("limit")
	if !ok {
		limit = "10"
	}

	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid offset"})
		return
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid limit"})
		return
	}

	userId := ctx.GetString("id")

	stmt := table.Organizations.
		INNER_JOIN(
			table.Users,
			table.Users.ID.EQ(table.Organizations.OwnerID),
		).
		SELECT(
			table.Organizations.AllColumns,
			table.Users.ID,
			table.Users.FirstName,
			table.Users.LastName,
			table.Users.Email,
			table.Users.Username,
		).
		WHERE(table.Organizations.OwnerID.EQ(jet.UUID(uuid.MustParse(userId)))).
		LIMIT(int64(limitInt)).
		OFFSET(int64(offsetInt))

	organizations := []organizationDao.OrganizationsWithOwner{}

	err = stmt.Query(db.DB, &organizations)

	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"data": organizations})
}
