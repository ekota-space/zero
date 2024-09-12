package organizationRoutes

import (
	organizationDao "github.com/ekota-space/zero/pkgs/organizations/dao"
	organizationModels "github.com/ekota-space/zero/pkgs/organizations/models"
	"github.com/ekota-space/zero/pkgs/root/db"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func PostCreate(ctx *gin.Context) {
	body := organizationDao.OrganizationInput{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	payload := organizationModels.Organizations{
		Name:        body.Name,
		OwnerID:     uuid.MustParse(body.OwnerID),
		Description: &body.Description,
	}

	result := db.DB.
		Create(&payload)

	if result.Error != nil {
		ctx.JSON(400, gin.H{"error": result.Error.Error()})
		return
	}

	payload.Owner = nil

	ctx.JSON(200, gin.H{"data": payload})
}
