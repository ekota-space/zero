package organizationRoutes

import (
	organizationModels "github.com/ekota-space/zero/pkgs/organizations/models"
	"github.com/ekota-space/zero/pkgs/root/db"
	"github.com/gin-gonic/gin"
)

func GetOrganization(ctx *gin.Context) {
	orgId := ctx.Param("id")

	organization := organizationModels.Organizations{}

	result := db.DB.Model(&organization).Where("id = ?", orgId).First(&organization)

	if result.Error != nil {
		ctx.JSON(404, gin.H{"error": "Organization not found"})
		return
	}

	ctx.JSON(200, gin.H{"data": organization})
}
