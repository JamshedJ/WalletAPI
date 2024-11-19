package gormRepo

import (
	"github.com/JamshedJ/WalletAPI/domain/repository"
	"gorm.io/gorm"
)

var _ repository.TransactionManagerI[*gorm.DB] = (*GormTxManager)(nil)

type GormTxManager struct {
	DB *gorm.DB
}

// Begin implements repository.TransactionManagerI.
func (d *GormTxManager) Begin() *gorm.DB {
	return d.DB.Begin()
}

// Commit implements repository.TransactionManagerI.
func (d *GormTxManager) Commit(db *gorm.DB) error {
	return db.Commit().Error
}

// Rollback implements repository.TransactionManagerI.
func (d *GormTxManager) Rollback(db *gorm.DB) error {
	return db.Rollback().Error
}
