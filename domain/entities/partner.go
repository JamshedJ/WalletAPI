package entities

import (
	"time"

	"github.com/google/uuid"
)

type Partner struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	SecretKey string    `json:"secret_key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
