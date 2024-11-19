package dto

import "time"

type GetTransactionsIn struct {
	UserID       uint
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
