package organizationRoutes

import (
	"github.com/ekota-space/zero/pkgs/common"
	organizationDao "github.com/ekota-space/zero/pkgs/organizations/dao"
	"github.com/ekota-space/zero/pkgs/response"
	"github.com/ekota-space/zero/pkgs/root/db/zero/public/model"
	"github.com/ekota-space/zero/pkgs/root/db/zero/public/table"
	"github.com/ekota-space/zero/pkgs/root/ql"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// @Summary		Create organization
// @Description	Create organization route
// @Tags			Organizations
// @Accept			json
// @Produce		json
// @Param			body	body		organizationDao.OrganizationInput	true	"Organization creation"
// @Success		201		{object}	response.SuccessDataResponse[model.Organizations]	"Organization created"
// @Failure		400		{object}	response.ErrorResponse[string]	"Invalid request"
// @Failure		403		{object}	response.ErrorResponse[string]	"Owner ID must be the same as the authenticated user"
// @Failure		500		{object}	response.ErrorResponse[string]	"Internal server error"
// @Router			/organizations [post]
func PostCreate(ctx *fiber.Ctx) error {
	body := organizationDao.OrganizationInput{}

	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(400).JSON(response.Error(err.Error()))
	}

	userId := ctx.Locals("id").(string)

	if userId != body.OwnerID {
		return ctx.Status(403).JSON(response.Error("owner_id must be the same as the authenticated user"))

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
		return ctx.Status(500).JSON(response.Error("Failed to start transaction"))
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
			return ctx.Status(400).JSON(response.Error("Slug already exists"))
		}

		return ctx.Status(400).JSON(response.Error(err.Error()))
	}

	tx.Commit()

	return ctx.Status(201).JSON(response.Success(result))
}
