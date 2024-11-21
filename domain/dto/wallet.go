package dto

import (
	"github.com/JamshedJ/WalletAPI/domain/errs"
	"github.com/google/uuid"
)

type TopUpWalletIn struct {
	PartnerID uuid.UUID
	Account   string
	Amount    float64
}

func (i *TopUpWalletIn) Validate() error {
	if i.Amount < 0 {
		return errs.ErrInvalidInput
	}
	if i.Account == "" {
		return errs.ErrInvalidInput
	}
	return nil
}
