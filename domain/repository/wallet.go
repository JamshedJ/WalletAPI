package repository

import (
	"context"

	"github.com/JamshedJ/WalletAPI/domain/dto"
	"github.com/JamshedJ/WalletAPI/domain/entities"
)

type WalletRepositoryI interface {
	Conn() any
	ExecuteTransaction(ctx context.Context, fn func(conn any) error) error
	UpdateWalletBalance(ctx context.Context, conn any, wallet *entities.Wallet) error
	GetWalletBalance(ctx context.Context, conn any, userID uint) (*entities.Wallet, error)
	CheckWalletExists(ctx context.Context, conn any, userID uint) (bool, error)
	CreateTransaction(ctx context.Context, conn any, transaction *entities.Transaction) error
	GetTransactions(ctx context.Context, conn any, params *dto.GetTransactionsIn) ([]*entities.Transaction, error)
}
