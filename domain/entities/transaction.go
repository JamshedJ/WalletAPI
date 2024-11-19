package entities

import "time"

type Transaction struct {
	ID        uint      `json:"id"`
	WalletID  uint      `json:"wallet_id"`
	UserID    uint      `json:"user_id"`
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}
