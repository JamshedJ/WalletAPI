package services

import (
	"context"

	"github.com/JamshedJ/WalletAPI/domain/dto"
	"github.com/JamshedJ/WalletAPI/domain/entities"
	"github.com/JamshedJ/WalletAPI/domain/errs"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func (s *ServicesTestSuite) TestGetWalletBalance_OK() {
	ctx := context.Background()
	userID := uint(1)
	expectedWallet := &entities.Wallet{
		ID:           10,
		UserID:       userID,
		Balance:      5000,
		IsIdentified: true,
	}

	s.WalletRepoMock.EXPECT().Conn().Return(mock.Anything)
	s.WalletRepoMock.EXPECT().GetWalletBalance(mock.Anything, s.WalletRepoMock.Conn(), userID).Return(expectedWallet, nil)

	wallet, err := s.Services.Wallet.GetWalletBalance(ctx, userID)

	require.NoError(s.T(), err)
	require.Equal(s.T(), wallet, expectedWallet)
	require.NotNil(s.T(), wallet)
}

func (s *ServicesTestSuite) TestTopUpWallet_OK() {
	ctx := context.Background()
	userID := uint(1)
	amount := 1000.0
	input := &dto.TopUpWalletIn{Amount: amount}
	existingWallet := &entities.Wallet{
		ID:           10,
		UserID:       userID,
		Balance:      5000,
		IsIdentified: true,
	}

	s.WalletRepoMock.EXPECT().ExecuteTransaction(ctx, mock.AnythingOfType("func(interface {}) error")).
		RunAndReturn(func(ctx context.Context, fn func(conn any) error) error {
			s.WalletRepoMock.EXPECT().CheckWalletExists(ctx, s.WalletRepoMock.Conn(), userID).Return(true, nil)
			s.WalletRepoMock.EXPECT().GetWalletBalance(ctx, s.WalletRepoMock.Conn(), userID).Return(existingWallet, nil)
			s.WalletRepoMock.EXPECT().UpdateWalletBalance(ctx, s.WalletRepoMock.Conn(), &entities.Wallet{
				ID:      existingWallet.ID,
				UserID:  existingWallet.UserID,
				Balance: existingWallet.Balance + amount,
			}).Return(nil)
			s.WalletRepoMock.EXPECT().CreateTransaction(ctx, s.WalletRepoMock.Conn(), &entities.Transaction{
				WalletID: existingWallet.ID,
				UserID:   userID,
				Amount:   amount,
			}).Return(nil)
			return fn(s.WalletRepoMock.Conn())
		}).Return(nil)

	err := s.Services.Wallet.TopUpWallet(ctx, userID, input)

	require.NoError(s.T(), err)
}

func (s *ServicesTestSuite) TestTopUpWallet_ValidationFailed() {
	ctx := context.Background()
	userID := uint(1)
	input := &dto.TopUpWalletIn{Amount: -1000}

	err := s.Services.Wallet.TopUpWallet(ctx, userID, input)

	require.Error(s.T(), err)
	require.ErrorIs(s.T(), err, errs.ErrValidationFailed)
}

func (s *ServicesTestSuite) TestGetMonthlySummary_OK() {
	ctx := context.Background()
	userID := uint(1)

	expectedTransactions := []*entities.Transaction{
		{Amount: 100.0},
		{Amount: 200.0},
	}

	s.WalletRepoMock.EXPECT().Conn().Return(mock.Anything)
	s.WalletRepoMock.EXPECT().GetTransactions(ctx, s.WalletRepoMock.Conn(), mock.Anything).Return(expectedTransactions, nil)

	summary, err := s.Services.Wallet.GetMonthlySummary(ctx, userID)

	require.NoError(s.T(), err)
	require.Equal(s.T(), len(expectedTransactions), summary.TotalTransactions)
	require.Equal(s.T(), 300.0, summary.TotalAmount)
}
