package auth

import (
	"time"

	"github.com/google/uuid"
)

type Users struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid();uniqueIndex" json:"id"`
	FirstName string    `gorm:"size:255" json:"first_name"`
	LastName  string    `gorm:"size:255" json:"last_name"`
	Username  string    `gorm:"size:255" json:"username"`
	Email     string    `gorm:"type:text;uniqueIndex" json:"email"`

	Password *string `gorm:"type:text" json:"password"`

	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	VerifiedAt      *time.Time `json:"verified_at"`
	PasswordResetAt *time.Time `json:"password_reset_at"`
}
