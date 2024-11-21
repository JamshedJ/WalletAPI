package gormRepo

import (
	"context"
	"errors"
	"time"

	"github.com/JamshedJ/WalletAPI/domain/dto"
	"github.com/JamshedJ/WalletAPI/domain/entities"
	"github.com/JamshedJ/WalletAPI/domain/errs"
	"github.com/JamshedJ/WalletAPI/domain/repository"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type gormWallet struct {
	ID           uint    `gorm:"primaryKey"`
	UserID       uint    `gorm:"not null"`
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
	e.Balance = w.Balance
	e.IsIdentified = w.IsIdentified
	e.CreatedAt = w.CreatedAt
	e.UpdatedAt = w.UpdatedAt
}

func (w *gormWallet) ToEntity() *entities.Wallet {
	return &entities.Wallet{
		ID:           w.ID,
		Balance:      w.Balance,
		IsIdentified: w.IsIdentified,
		CreatedAt:    w.CreatedAt,
		UpdatedAt:    w.UpdatedAt,
	}
}

var _ repository.WalletRepositoryI = (*GormWalletRepo)(nil)

type GormWalletRepo struct {
}

func (g *GormWalletRepo) Conn() any {
	return DB
}

func (g *GormWalletRepo) ExecuteTransaction(ctx context.Context, fn func(conn any) error) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		return fn(tx)
	})
}

func (g *GormWalletRepo) GetWalletBalance(ctx context.Context, conn any, account string) (*entities.Wallet, error) {
	db := conn.(*gorm.DB)
	var gw = &gormWallet{}

	err := db.WithContext(ctx).Where("account = ?", account).First(&gw).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.ErrWalletDoesNotExist
		}
		return nil, err
	}

	return gw.ToEntity(), nil
}

func (g *GormWalletRepo) CheckWalletExists(ctx context.Context, conn any, account string) (bool, error) {
	db := conn.(*gorm.DB)
	var gw = &gormWallet{}

	res := db.WithContext(ctx).Model(&gw).Where("account = ?", account).First(&gw)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, res.Error
	}

	return true, nil
}

func (g *GormWalletRepo) UpdateWalletBalance(ctx context.Context, conn any, wallet *entities.Wallet) error {
	db := conn.(*gorm.DB)
	result := db.WithContext(ctx).Model(&gormWallet{}).Where("id = ?", wallet.ID).Updates(map[string]any{
		"balance": wallet.Balance,
	})
	if result.Error != nil {
		return result.Error
	}

	return nil
}

type gormTransaction struct {
	ID        uint    `gorm:"primaryKey"`
	WalletID  uint    `gorm:"not null"`
	PartnerID uuid.UUID    `gorm:"not null"`
	Amount    float64 `gorm:"not null"`
	CreatedAt time.Time
}

func (gormTransaction) TableName() string {
	return "transactions"
}

func (t *gormTransaction) ParseEntity(e *entities.Transaction) {
	e.ID = t.ID
	e.WalletID = t.WalletID
	e.PartnerID = t.PartnerID
	e.Amount = t.Amount
	e.CreatedAt = t.CreatedAt
}

func (t *gormTransaction) ToEntity() *entities.Transaction {
	return &entities.Transaction{
		ID:        t.ID,
		WalletID:  t.WalletID,
		PartnerID:    t.PartnerID,
		Amount:    t.Amount,
		CreatedAt: t.CreatedAt,
	}
}

func (g *GormWalletRepo) CreateTransaction(ctx context.Context, conn any, transaction *entities.Transaction) error {
	db := conn.(*gorm.DB)
	err := db.WithContext(ctx).Create(transaction).Error
	if err != nil {
		return err
	}

	return nil
}

func (g *GormWalletRepo) GetTransactions(ctx context.Context, conn any, params *dto.GetTransactionsIn) ([]*entities.Transaction, error) {
	db := conn.(*gorm.DB)
	var gt = []*gormTransaction{}

	err := db.WithContext(ctx).Where("partner_id = ?", params.PartnerID).Find(&gt).Error
	if err != nil {
		return nil, err
	}

	transactions := []*entities.Transaction{}
	for _, t := range gt {
		transaction := t.ToEntity()
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}
