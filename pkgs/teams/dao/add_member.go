package teamsDao

type AddMemberInput struct {
	ID string `json:"id" binding:"required"`
}
