package organizationRoutes

import (
	"github.com/ekota-space/zero/pkgs/common"
	organizationDao "github.com/ekota-space/zero/pkgs/organizations/dao"
	"github.com/ekota-space/zero/pkgs/response"
	"github.com/ekota-space/zero/pkgs/root/db/zero/public/model"
	"github.com/ekota-space/zero/pkgs/root/db/zero/public/table"
	"github.com/ekota-space/zero/pkgs/root/ql"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//	@Summary		Create organization
//	@Description	Create organization route
//	@Tags			Organizations
//	@Accept			json
//	@Produce		json
//	@Param			body	body		response.SuccessDataResponse[organizationDao.OrganizationInput]	true	"Organization creation"
//	@Success		201		{object}	model.Organizations												"Organization created"
//	@Failure		400		{object}	response.ErrorResponse[string]									"Invalid request"
//	@Failure		403		{object}	response.ErrorResponse[string]									"Owner ID must be the same as the authenticated user"
//	@Failure		500		{object}	response.ErrorResponse[string]									"Internal server error"
//	@Router			/organizations [post]
func PostCreate(ctx *gin.Context) {
	body := organizationDao.OrganizationInput{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(400, response.Error(err.Error()))
		return
	}

	userId := ctx.GetString("id")

	if userId != body.OwnerID {
		ctx.JSON(403, response.Error("owner_id must be the same as the authenticated user"))
		return
	}

	ownerId := uuid.MustParse(body.OwnerID)
	payload := model.Organizations{
		Name:        body.Name,
		OwnerID:     ownerId,
		Description: &body.Description,
		Slug:        body.Slug,
	}

	tx, err := ql.GetDB().Begin()

	if err != nil {
		ctx.JSON(500, response.Error("Failed to start transaction"))
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
			ctx.JSON(400, response.Error("Slug already exists"))
			return
		}

		ctx.JSON(400, response.Error(err.Error()))
		return
	}

	tx.Commit()

	ctx.JSON(201, response.Success(result))
}
