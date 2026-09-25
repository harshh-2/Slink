package models

import (
	"time"
	"github.com/google/uuid"
)

type Link struct {
	ID uuid.UUID
	UserID uuid.UUID 
	ShortCode string 
	LongURL string 
	Title *string 
	ClickCount int64 
	LastClickedAt *time.Time 
	ExpiresAt *time.Time 
	PasswordHash *string
	IsActive bool 
	CreatedAt time.Time 
	UpdatedAt time.Time 
	DeletedAt *time.Time 
}