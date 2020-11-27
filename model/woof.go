package model

import "time"

// Woof db struct
type Woof struct {
	ID        string    `db:"id"`
	Body      string    `db:"body"`
	CreatedAt time.Time `db:"created_at"`
}

// WoofRequest json struct
type WoofRequest struct {
	Message string `json:"message"`
}

// WoofResponse json struct
type WoofResponse struct {
	ID string `json:"id"`
}
