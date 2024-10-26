package organizationRoutes

import (
	"github.com/ekota-space/zero/pkgs/response"
	"github.com/ekota-space/zero/pkgs/root/db/zero/public/model"
	"github.com/ekota-space/zero/pkgs/root/db/zero/public/table"
	"github.com/ekota-space/zero/pkgs/root/ql"
	"github.com/gin-gonic/gin"
	jet "github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"
)

//	@Summary		List organizations
//	@Description	List organizations route
//	@Tags			Organizations
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	response.SuccessDataResponse[[]model.Organizations]	"List of organizations"
//	@Failure		400	{object}	response.ErrorResponse[string]						"Invalid request"
//	@Router			/organizations [get]
func GetList(ctx *gin.Context) {
	userId := ctx.GetString("id")

	stmt := table.Organizations.
		LEFT_JOIN(
			table.OrganizationMembers,
			table.OrganizationMembers.OrganizationID.EQ(table.Organizations.ID).
				AND(table.OrganizationMembers.ID.EQ(jet.UUID(uuid.MustParse(userId)))),
		).
		SELECT(table.Organizations.AllColumns).
		WHERE(
			table.Organizations.OwnerID.EQ(jet.UUID(uuid.MustParse(userId))).
				OR(table.OrganizationMembers.ID.EQ(jet.UUID(uuid.MustParse(userId)))),
		)

	organizations := []model.Organizations{}

	err := stmt.Query(ql.GetDB(), &organizations)

	if err != nil {
		ctx.JSON(400, response.Error(err.Error()))
		return
	}

	ctx.JSON(200, response.Success(organizations))
}
