package organizationRoutes

import (
	"strconv"

	organizationModels "github.com/ekota-space/zero/pkgs/organizations/models"
	"github.com/ekota-space/zero/pkgs/root/db"
	"github.com/gin-gonic/gin"
)

func GetList(ctx *gin.Context) {
	organizations := []organizationModels.Organizations{}

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

	result := db.DB.
		Table("organizations").
		Select("organizations.*, owner.email, owner.username, owner.first_name, owner.last_name").
		Joins("left join users as owner on organizations.owner_id = owner.id").
		Offset(offsetInt).
		Limit(limitInt).
		Find(&organizations)

	if result.Error != nil {
		ctx.JSON(400, gin.H{"error": result.Error.Error()})
		return
	}

	ctx.JSON(200, gin.H{"data": organizations})
}
