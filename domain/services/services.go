package services

import (
	"github.com/JamshedJ/WalletAPI/domain/repository"
	"github.com/rs/zerolog"
)

// ServiceFacade объединяет все сервисы приложения в единую точку доступа.
// Этот фасад упрощает взаимодействие между хендлерами и сервисами,
// предоставляя унифицированный интерфейс для работы с логикой бизнес-процессов
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
