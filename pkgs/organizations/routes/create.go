package organizationRoutes

import (
	organizationDao "github.com/ekota-space/zero/pkgs/organizations/dao"
	"github.com/ekota-space/zero/pkgs/root/db"
	"github.com/ekota-space/zero/pkgs/root/db/zero/public/model"
	"github.com/ekota-space/zero/pkgs/root/db/zero/public/table"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func PostCreate(ctx *gin.Context) {
	body := organizationDao.OrganizationInput{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ownerId := uuid.MustParse(body.OwnerID)
	payload := model.Organizations{
		Name:        body.Name,
		OwnerID:     ownerId,
		Description: &body.Description,
	}

	tx, err := db.DB.Begin()

	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to start transaction"})
		return
	}

	stmt := table.Organizations.INSERT(table.Organizations.Name, table.Organizations.OwnerID, table.Organizations.Description).MODEL(payload).RETURNING(table.Organizations.AllColumns)

	result := model.Organizations{}
	err = stmt.Query(tx, &result)

	if err != nil {
		tx.Rollback()

		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"data": result})
}
