package dto

import "time"

type GetTransactionsIn struct {
	UserID       string
	StartOfMonth time.Time
	EndOfMonth   time.Time
}

type CreateTransactionIn struct {
	WalletID uint
	Amount   float64
}

type GetMonthlySummaryOut struct {
	TotalTransactions int
	TotalAmount       float64
}
