package teamsDao

type CreateTeamInput struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Slug        string `json:"slug" binding:"required"`
}
