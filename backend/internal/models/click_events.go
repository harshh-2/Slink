package models

import (
	"time"
	"github.com/google/uuid"
)

type ClickEvent struct {
	ID  uuid.UUID `json:"id"`
	LinkID uuid.UUID `json:"link_id"`
	Browser string `json:"browser"`
	OS string `json:"os"`
	Device string `json:"device"`
	Referer string `json:"referer"`
	IPHash string `json:"-"`
	ClickedAt time.Time `json:"clicked_at"`
}