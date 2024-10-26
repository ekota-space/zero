package organizations

import (
	"github.com/ekota-space/zero/pkgs/response"
	"github.com/gin-gonic/gin"
)

func AccessCheckMiddleware(accessRole string) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		userId := ctx.GetString("id")
		orgSlug := ctx.Param("orgSlug")

		accessLevel, org, err := GetAccessLevel(userId, orgSlug)

		if err != nil {
			ctx.JSON(500, response.Error("internal server error"))
			ctx.Abort()
			return
		}
		// Owner can access everything
		isOwner := accessLevel == OWNER && (accessRole == ADMIN || accessRole == MEMBER)
		// Admin can access everything except owner's
		isAdmin := accessLevel == ADMIN && accessRole == MEMBER

		ctx.Set("organizationId", org.ID.String())

		if accessLevel == accessRole || isOwner || isAdmin {
			ctx.Next()
			return
		}

		ctx.JSON(403, response.Error("forbidden"))
		ctx.Abort()
	}
}
