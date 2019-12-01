package model

import "time"

// Base contains common fields for all entity models
type Base struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	//DeletedAt time.Time `json:"deleted_at,omitempty" pg:",soft_delete"`
}
