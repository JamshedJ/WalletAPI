package entities

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID        uint      `json:"id"`
	WalletID  uint      `json:"wallet_id"`
	PartnerID uuid.UUID `json:"partner_id"`
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}
