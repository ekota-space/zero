package organizationDao

type OrganizationInput struct {
	Name    string `json:"name" binding:"required"`
	OwnerID string `json:"owner_id" binding:"required"`

	Description string `json:"description"`
}
