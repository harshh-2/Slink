package models

import (
	"time"
	"github.com/google/uuid"
)

type APIKey struct {
	ID uuid.UUID 
	UserID uuid.UUID 
	Name string
	KeyHash string 
	LastUsedAt *time.Time 
	ExpiresAt *time.Time 
	RevokedAt *time.Time 
	CreatedAt time.Time 
}