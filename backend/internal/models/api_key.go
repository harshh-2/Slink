package models

import (
	"time"
	"github.com/google/uuid"
)

type APIKey struct {
	ID uuid.UUID `json:"id"`
	UserID uuid.UUID `json:"user_id"`
	Name string `json:"name"`
	KeyHash string `json:"-"`
	LastUsedAt *time.Time `json:"last_used_at"`
	ExpiresAt *time.Time `json:"expires_at"`
	RevokedAt *time.Time `json:"revoked_at"`
	CreatedAt time.Time `json:"created_at"`
}