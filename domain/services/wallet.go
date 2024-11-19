package services

import (
	"context"
	"time"

	"github.com/JamshedJ/WalletAPI/domain/dto"
	"github.com/JamshedJ/WalletAPI/domain/entities"
	"github.com/JamshedJ/WalletAPI/domain/errs"
	"github.com/JamshedJ/WalletAPI/domain/repository"
	"github.com/rs/zerolog"
)

type WalletServiceI interface {
	TopUpWallet(ctx context.Context, userID string, in *dto.TopUpWalletIn) error
	GetWalletBalance(ctx context.Context, userID string) (*entities.Wallet, error)
	CheckWalletExists(ctx context.Context, userID string) (bool, error)
	GetMonthlySummary(ctx context.Context, userID string) (*dto.GetMonthlySummaryOut, error)
}

var _ WalletServiceI = (*WalletService)(nil)

type WalletService struct {
	WalletRepo      repository.WalletRepositoryI
	Logger          zerolog.Logger
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

func (s *WalletService) CheckWalletExists(ctx context.Context, userID string) (bool, error) {
	logger := s.Logger.With().Str("userID", userID).Logger()
	isWalletExists, err := s.WalletRepo.CheckWalletExists(ctx, s.WalletRepo.Conn(), userID)
	if err != nil {
		logger.Error().Err(err).Msg("failed to check wallet existence")
		return false, err
	}

	return isWalletExists, nil
}

func (s *WalletService) TopUpWallet(ctx context.Context, userID string, in *dto.TopUpWalletIn) error {
	logger := s.Logger.With().Str("userID", userID).Float64("amount", in.Amount).Logger()

	isWalletExists, err := s.WalletRepo.CheckWalletExists(ctx, s.WalletRepo.Conn(), userID)
	if err != nil {
		logger.Error().Err(err).Msg("failed to check wallet existence")
		return err
	}

	if !isWalletExists {
		logger.Error().Msg("wallet does not exist")
		return errs.ErrWalletDoesNotExist
	}

	wallet, err := s.WalletRepo.GetWalletBalance(ctx, s.WalletRepo.Conn(), userID)
	if err != nil {
		logger.Error().Err(err).Msg("failed to get wallet balance")
		return err
	}

	if wallet.Balance+in.Amount > 100000 {
		logger.Error().Msg("wallet balance cannot exceed 100000")
		return errs.ErrWalletBalanceLimitExceeded
	}

	err = s.WalletRepo.UpdateWallet(ctx, s.WalletRepo.Conn(), &entities.Wallet{
		ID: wallet.ID,
		UserID:  userID,
		Balance: wallet.Balance + in.Amount,
	})
	if err != nil {
		logger.Error().Err(err).Msg("failed to update wallet balance")
		return err
	}

	err = s.WalletRepo.CreateTransaction(ctx, s.WalletRepo.Conn(), &entities.Transaction{
		WalletID: wallet.ID,
		Amount:   in.Amount,
	})
	if err != nil {
		logger.Error().Err(err).Msg("failed to create transaction")
		return err
	}

	return nil
}

func (s *WalletService) GetMonthlySummary(ctx context.Context, userID string) (*dto.GetMonthlySummaryOut, error) {
	logger := s.Logger.With().Str("userID", userID).Logger()

	startOfMonth := time.Now().AddDate(0, 0, -time.Now().Day()+1)
	endOfMonth := startOfMonth.AddDate(0, 1, 0).Add(-time.Second)

	transactions, err := s.WalletRepo.GetTransactions(ctx, s.WalletRepo.Conn(), &dto.GetTransactionsIn{
		UserID:       userID,
		StartOfMonth: startOfMonth,
		EndOfMonth:   endOfMonth,
	})
	if err != nil {
		logger.Error().Err(err).Msg("failed to get transactions")
		return nil, err
	}

	var totalAmount float64
	for _, transaction := range transactions {
		totalAmount += transaction.Amount
	}

	var summary = &dto.GetMonthlySummaryOut{
		TotalTransactions: len(transactions),
		TotalAmount:       totalAmount,
	}
	return summary, nil
}
