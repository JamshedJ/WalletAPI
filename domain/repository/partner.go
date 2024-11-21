package repository

import (
	"context"

	"github.com/JamshedJ/WalletAPI/domain/entities"
)

type PartnerRepositoryI interface {
	Conn() any
	GetPartnerByID(ctx context.Context, id string) (*entities.Partner, error)
}
