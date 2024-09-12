package organizationDao

import (
	"github.com/ekota-space/zero/pkgs/root/db/zero/public/model"
	"github.com/google/uuid"
)

type OrganizationsWithOwner struct {
	model.Organizations
	Owner struct {
		ID        uuid.UUID `sql:"primary_key" json:"id"`
		FirstName string    `json:"first_name"`
		LastName  string    `json:"last_name"`
		Username  string    `json:"username"`
		Email     string    `json:"email"`
	} `alias:"users.*" json:"owner"` // Alias is required for go-jet to map the joined fields correctly
}
