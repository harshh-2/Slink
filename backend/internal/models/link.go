package models

import (
	"time"
	"github.com/google/uuid"
)

type Link struct {
	ID uuid.UUID `json:"id"`
	UserID uuid.UUID `json:"user_id"`
	ShortCode string `json:"short_code"`
	LongURL string `json:"long_url"`
	Title string `json:"title"`
	ClickCount int64 `json:"click_count"`
	LastClickedAt *time.Time `json:"last_clicked_at"`
	ExpiresAt *time.Time `json:"expires_at"`
	PasswordHash *string `json:"-"`
	IsActive bool `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}