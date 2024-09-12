package organizationModels

import (
	"time"

	authModels "github.com/ekota-space/zero/pkgs/auth/models"
	"github.com/google/uuid"
)

type Organizations struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid();uniqueIndex" json:"id"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at"`

	Name        string  `gorm:"size:255;not null" json:"name"`
	Description *string `gorm:"size:255;" json:"description"`

	OwnerID uuid.UUID         `json:"owner_id"`
	Owner   *authModels.Users `json:"owner,omitempty"`
}
