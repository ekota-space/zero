package authModels

import (
	"time"

	"github.com/google/uuid"
)

type Users struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid();uniqueIndex" json:"id"`
	FirstName string    `gorm:"size:255;not null" json:"first_name"`
	LastName  string    `gorm:"size:255;not null" json:"last_name"`
	Username  string    `gorm:"size:16;not null;uniqueIndex" json:"username"`
	Email     string    `gorm:"type:text;not null;uniqueIndex" json:"email"`

	Password *string `gorm:"type:text" json:"password"`

	CreatedAt       time.Time  `gorm:"not null" json:"created_at"`
	UpdatedAt       time.Time  `gorm:"not null" json:"updated_at"`
	VerifiedAt      *time.Time `json:"verified_at"`
	PasswordResetAt *time.Time `json:"password_reset_at"`
}
