package organizationRoutes

import (
	"github.com/ekota-space/zero/pkgs/common"
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

	userId := ctx.GetString("id")

	if userId != body.OwnerID {
		ctx.JSON(403, gin.H{"error": "owner_id must be the same as the authenticated user"})
		return
	}

	ownerId := uuid.MustParse(body.OwnerID)
	payload := model.Organizations{
		Name:        body.Name,
		OwnerID:     ownerId,
		Description: &body.Description,
		Slug:        body.Slug,
	}

	tx, err := db.DB.Begin()

	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to start transaction"})
		return
	}

	stmt := table.Organizations.INSERT(
		table.Organizations.Name,
		table.Organizations.OwnerID,
		table.Organizations.Description,
		table.Organizations.Slug,
	).
		MODEL(payload).
		RETURNING(table.Organizations.AllColumns)

	result := model.Organizations{}
	err = stmt.Query(tx, &result)

	if err != nil {
		tx.Rollback()

		if common.IsDuplicateKeyError(err) {
			ctx.JSON(400, gin.H{"error": "Slug already exists"})
			return
		}

		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	tx.Commit()

	ctx.JSON(201, gin.H{"data": result})
}
