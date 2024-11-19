package services

import (
	"github.com/JamshedJ/WalletAPI/domain/repository"
	"github.com/rs/zerolog"
)

type ServiceFacade struct {
	Wallet WalletServiceI
}

func NewServiceFacade(logger zerolog.Logger, repo repository.RepositoryFacade) *ServiceFacade {
	return &ServiceFacade{
		Wallet: &WalletService{
			WalletRepo: repo.WalletRepositoryI,
			Logger:     logger.With().Str("service", "wallet").Logger(),
		},
	}
}
