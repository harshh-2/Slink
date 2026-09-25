package models

import (
	"time"
	"github.com/google/uuid"
)

type User struct {
	ID uuid.UUID 
	ClerkID string    
	Username string
	Email string    
	IsVerified bool      
	CreatedAt time.Time 
	UpdatedAt time.Time  
	DeletedAt *time.Time
}