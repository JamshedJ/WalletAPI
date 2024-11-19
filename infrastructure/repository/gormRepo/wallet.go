package gormRepo

import (
	"context"
	"time"

	"github.com/JamshedJ/WalletAPI/domain/entities"
)

type gormWallet struct {
	ID           uint    `gorm:"primaryKey"`
	UserID       string  `gorm:"not null"`
	Balance      float64 `gorm:"not null"`
	IsIdentified bool    `gorm:"not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (gormWallet) TableName() string {
	return "wallets"
}

func (w *gormWallet) ParseEntity(e *entities.Wallet) {
	e.ID = w.ID
	e.UserID = w.UserID
	e.Balance = w.Balance
	e.IsIdentified = w.IsIdentified
	e.CreatedAt = w.CreatedAt
	e.UpdatedAt = w.UpdatedAt
}

func (w *gormWallet) ToEntity() *entities.Wallet {
	return &entities.Wallet{
		ID:           w.ID,
		UserID:       w.UserID,
		Balance:      w.Balance,
		IsIdentified: w.IsIdentified,
		CreatedAt:    w.CreatedAt,
		UpdatedAt:    w.UpdatedAt,
	}
}

type GormWalletRepo struct {
}

func (g *GormWalletRepo) Conn() any {
	return DB
}

func (g *GormWalletRepo) GetWalletBalance(ctx context.Context, conn any, userID string) (*entities.Wallet, error) {
	var gw = &gormWallet{}

	err := DB.WithContext(ctx).Where("user_id = ?", userID).First(&gw).Error
	if err != nil {
		return nil, err
	}

	return gw.ToEntity(), nil
}
