package services

import (
	"context"
	"errors"
	"time"

	"github.com/JamshedJ/WalletAPI/domain/dto"
	"github.com/JamshedJ/WalletAPI/domain/entities"
	"github.com/JamshedJ/WalletAPI/domain/errs"
	"github.com/JamshedJ/WalletAPI/domain/repository"
	"github.com/rs/zerolog"
)

type WalletServiceI interface {
	GetWalletBalance(ctx context.Context, userID uint) (*entities.Wallet, error)
	CheckWalletExists(ctx context.Context, userID uint) (bool, error)
	TopUpWallet(ctx context.Context, userID uint, in *dto.TopUpWalletIn) error
	GetMonthlySummary(ctx context.Context, userID uint) (*dto.GetMonthlySummaryOut, error)
}

var _ WalletServiceI = (*WalletService)(nil)

type WalletService struct {
	WalletRepo repository.WalletRepositoryI
	Logger     zerolog.Logger
}

func (s *WalletService) GetWalletBalance(ctx context.Context, userID uint) (*entities.Wallet, error) {
	logger := s.Logger.With().Uint("userID", userID).Logger()
	balance, err := s.WalletRepo.GetWalletBalance(ctx, s.WalletRepo.Conn(), userID)
	if err != nil {
		logger.Error().Err(err).Msg("failed to get wallet balance")
		return nil, err
	}
	return balance, nil
}

func (s *WalletService) CheckWalletExists(ctx context.Context, userID uint) (bool, error) {
	logger := s.Logger.With().Uint("userID", userID).Logger()
	isWalletExists, err := s.WalletRepo.CheckWalletExists(ctx, s.WalletRepo.Conn(), userID)
	if err != nil {
		logger.Error().Err(err).Msg("failed to check wallet existence")
		return false, err
	}

	return isWalletExists, nil
}

func (s *WalletService) TopUpWallet(ctx context.Context, userID uint, in *dto.TopUpWalletIn) error {
	logger := s.Logger.With().Uint("userID", userID).Float64("amount", in.Amount).Logger()

	if err := in.Validate(); err != nil {
		logger.Error().Err(err).Msg("validation failed")
		return errors.Join(errs.ErrValidationFailed, err)
	}

	// В данной реализации используется транзакция, хотя она не является строго необходимой, 
	// поскольку логика не включает несколько зависимых операций, 
	// которые требуют атомарности. Транзакция добавлена для того, чтобы продемонстрировать, как она могла бы выглядеть.
	err := s.WalletRepo.ExecuteTransaction(ctx, func(conn any) error {
		isWalletExists, err := s.WalletRepo.CheckWalletExists(ctx, conn, userID)
		if err != nil {
			logger.Error().Err(err).Msg("failed to check wallet existence")
			return err
		}

		if !isWalletExists {
			logger.Error().Msg("wallet does not exist")
			return errs.ErrWalletDoesNotExist
		}

		receiverWallet, err := s.WalletRepo.GetWalletBalance(ctx, conn, in.UserID)
		if err != nil {
			logger.Error().Err(err).Msg("failed to get receiver's wallet balance")
			return err
		}

		switch {
		case receiverWallet.IsIdentified && receiverWallet.Balance+in.Amount > 100000:
			logger.Error().Msg("receiver's wallet balance exceeds the limit")
			return errs.ErrWalletBalanceLimitExceeded
		case !receiverWallet.IsIdentified && receiverWallet.Balance+in.Amount > 10000:
			logger.Error().Msg("receiver's wallet balance exceeds the limit")
			return errs.ErrWalletBalanceLimitExceeded
		}

		// Increase the balance of the receiver's wallet
		err = s.WalletRepo.UpdateWalletBalance(ctx, conn, &entities.Wallet{
			ID:      receiverWallet.ID,
			UserID:  receiverWallet.UserID,
			Balance: receiverWallet.Balance + in.Amount,
		})
		if err != nil {
			logger.Error().Err(err).Msg("failed to update receiver's wallet balance")
			return err
		}

		err = s.WalletRepo.CreateTransaction(ctx, conn, &entities.Transaction{
			WalletID: in.WalletID,
			UserID:   in.UserID,
			Amount:   in.Amount,
		})
		if err != nil {
			logger.Error().Err(err).Msg("failed to create transaction")
			return err
		}

		return nil
	})
	if err != nil {
		logger.Error().Err(err).Msg("failed to top up wallet")
		return err
	}

	logger.Info().Msg("wallet topped up successfully")
	return nil
}

func (s *WalletService) GetMonthlySummary(ctx context.Context, userID uint) (*dto.GetMonthlySummaryOut, error) {
	logger := s.Logger.With().Uint("userID", userID).Logger()

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
