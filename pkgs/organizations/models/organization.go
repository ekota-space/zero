package organizationModels

import (
	"time"

	"github.com/google/uuid"
)

type Organizations struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid();uniqueIndex" json:"id"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at"`

	Name string `gorm:"size:255;not null" json:"name"`
}
