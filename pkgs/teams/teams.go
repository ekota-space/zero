package teams

import (
	"github.com/ekota-space/zero/pkgs/root/db/zero/public/model"
	"github.com/google/uuid"
)

const (
	NONE   = "none"
	MEMBER = "member"
	ADMIN  = "admin"
)

func GetAccessLevel(userId string, teamSlug string, orgSlug string) (string, *model.Teams, error) {
	org, err := GetOrganizationBySlug(orgSlug)

	if err != nil {
		return NONE, &org, err
	}

	if org.OwnerID == uuid.MustParse(userId) {
		return OWNER, &org, nil
	}

	admin, err := GetOrganizationAdminByIds(userId, org.ID.String())

	if err != nil || admin.UserID == uuid.Nil {
		return NONE, &org, err
	}

	if admin.UserID == uuid.MustParse(userId) {
		return ADMIN, &org, nil
	}

	member, err := GetOrganizationMemberByIds(userId, org.ID.String())

	if err != nil || member.UserID == uuid.Nil {
		return NONE, &org, err
	}

	if member.UserID == uuid.MustParse(userId) {
		return MEMBER, &org, nil
	}

	return NONE, &org, nil
}
