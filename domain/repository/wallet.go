package repository

import (
	"context"

	"github.com/JamshedJ/WalletAPI/domain/entities"
)

type WalletRepositoryI interface {
	Conn() any
	GetWalletBalance(ctx context.Context, conn any, userID string) (*entities.Wallet, error)
}
