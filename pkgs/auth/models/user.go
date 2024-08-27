package auth

import (
	"time"

	"github.com/google/uuid"
)

type Users struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid();uniqueIndex"`
	FirstName string    `gorm:"size:255"`
	LastName  string    `gorm:"size:255"`
	Username  string    `gorm:"size:255"`
	Email     string    `gorm:"type:text;uniqueIndex"`

	Password *string `gorm:"type:text"`

	CreatedAt       time.Time
	UpdatedAt       time.Time
	VerifiedAt      *time.Time
	PasswordResetAt *time.Time
}
