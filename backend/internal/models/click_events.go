package models

import (
	"time"
	"github.com/google/uuid"
)

type ClickEvent struct {
	ID  uuid.UUID 
	LinkID uuid.UUID 
	Browser string 
	OS string 
	Device string 
	Referer string 
	IPHash string 
	ClickedAt time.Time 
}