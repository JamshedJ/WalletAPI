package services

import (
	"context"

	"github.com/JamshedJ/WalletAPI/domain/entities"
	"github.com/JamshedJ/WalletAPI/domain/repository"
	"github.com/rs/zerolog"
)

type WalletServiceI interface {
	GetWalletBalance(ctx context.Context, userID string) (*entities.Wallet, error)
}

type WalletService struct {
	WalletRepo repository.WalletRepositoryI
	Logger     zerolog.Logger
}

func (s *WalletService) GetWalletBalance(ctx context.Context, userID string) (*entities.Wallet, error) {
	logger := s.Logger.With().Str("userID", userID).Logger()
	balance, err := s.WalletRepo.GetWalletBalance(ctx, s.WalletRepo.Conn(), userID)
	if err != nil {
		logger.Error().Err(err).Msg("failed to get wallet balance")
		return nil, err
	}
	return balance, nil
}