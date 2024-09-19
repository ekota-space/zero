package organizations

import (
	"github.com/ekota-space/zero/pkgs/root/db/zero/public/model"
	"github.com/ekota-space/zero/pkgs/root/db/zero/public/table"
	"github.com/ekota-space/zero/pkgs/root/ql"
	jet "github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"
)

const (
	NONE   = "none"
	OWNER  = "owner"
	MEMBER = "member"
	ADMIN  = "admin"
)

func GetOrganizationBySlug(slug string) (model.Organizations, error) {
	stmt := table.Organizations.
		SELECT(
			table.Organizations.AllColumns,
		).
		WHERE(
			table.Organizations.Slug.EQ(jet.String(slug)),
		)

	var org model.Organizations

	err := stmt.Query(ql.GetDB(), &org)

	return org, err
}

func GetOrganizationMemberByIds(userId string, orgId string) (model.OrganizationMembers, error) {
	stmt := table.OrganizationMembers.
		SELECT(
			table.OrganizationMembers.AllColumns,
		).
		WHERE(
			table.OrganizationMembers.UserID.EQ(jet.UUID(uuid.MustParse(userId))).
				AND(
					table.OrganizationMembers.OrganizationID.EQ(jet.UUID(uuid.MustParse(orgId))),
				),
		)

	var member model.OrganizationMembers

	err := stmt.Query(ql.GetDB(), &member)

	return member, err
}

func GetOrganizationAdminByIds(userId string, orgId string) (model.OrganizationAdmins, error) {
	stmt := table.OrganizationAdmins.
		SELECT(
			table.OrganizationAdmins.AllColumns,
		).
		WHERE(
			table.OrganizationAdmins.UserID.EQ(jet.UUID(uuid.MustParse(userId))).
				AND(
					table.OrganizationAdmins.OrganizationID.EQ(jet.UUID(uuid.MustParse(orgId))),
				),
		)

	var admin model.OrganizationAdmins

	err := stmt.Query(ql.GetDB(), &admin)

	return admin, err

}

func GetAccessLevel(userId string, orgSlug string) (string, *model.Organizations, error) {
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
