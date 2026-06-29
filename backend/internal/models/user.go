package models

import (
	"time"
	"github.com/google/uuid"
)

type User struct {
	ID uuid.UUID `json:"id"`
	ClerkID string    `json:"clerk_id"`
	Username string    `json:"username"`
	Email string    `json:"email"`
	IsVerified bool      `json:"is_verified"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}