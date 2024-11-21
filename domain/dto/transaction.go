package dto

import (
	"time"

	"github.com/google/uuid"
)

type GetTransactionsIn struct {
	PartnerID    uuid.UUID
	StartOfMonth time.Time
	EndOfMonth   time.Time
}

type CreateTransactionIn struct {
	WalletID uint
	Amount   float64
}

type GetMonthlySummaryOut struct {
	TotalTransactions int     `json:"total_transactions"`
	TotalAmount       float64 `json:"total_amount"`
}
