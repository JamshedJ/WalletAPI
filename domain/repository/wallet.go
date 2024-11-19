package repository

import (
	"context"

	"github.com/JamshedJ/WalletAPI/domain/dto"
	"github.com/JamshedJ/WalletAPI/domain/entities"
)

type WalletRepositoryI interface {
	Conn() any
	UpdateWallet(ctx context.Context, conn any, wallet *entities.Wallet) error
	GetWalletBalance(ctx context.Context, conn any, userID string) (*entities.Wallet, error)
	CheckWalletExists(ctx context.Context, conn any, userID string) (bool, error)
	CreateTransaction(ctx context.Context, conn any, transaction *entities.Transaction) error
	GetTransactions(ctx context.Context, conn any, params *dto.GetTransactionsIn) ([]*entities.Transaction, error)
}
