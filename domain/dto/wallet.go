package dto

import "github.com/JamshedJ/WalletAPI/domain/errs"

type TopUpWalletIn struct {
	WalletID uint
	UserID   uint
	Amount   float64
}

func (i *TopUpWalletIn) Validate() error {
	if i.Amount < 0 {
		return errs.ErrInvalidInput
	}
	return nil
}
