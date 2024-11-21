package entities

import "time"

type Wallet struct {
	ID           uint      `json:"id"`
	Account      string    `json:"account"`
	Balance      float64   `json:"balance"`
	IsIdentified bool      `json:"is_identified"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
