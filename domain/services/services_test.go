package services

import (
	"os"
	"testing"

	"github.com/JamshedJ/WalletAPI/domain/repository"
	repoMock "github.com/JamshedJ/WalletAPI/mocks/repository"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/suite"
)

type ServicesTestSuite struct {
	suite.Suite
	WalletRepoMock *repoMock.MockWalletRepositoryI
	PartnerRepoMock *repoMock.MockPartnerRepositoryI
	Services       ServiceFacade
}

func (s *ServicesTestSuite) SetupSuite() {
	walletRepoMock := repoMock.NewMockWalletRepositoryI(s.T())
	partnerRepoMock := repoMock.NewMockPartnerRepositoryI(s.T())
	s.WalletRepoMock = walletRepoMock
	s.PartnerRepoMock = partnerRepoMock
	repoFacade := repository.RepositoryFacade{
		WalletRepositoryI: walletRepoMock,
		PartnerRepositoryI: partnerRepoMock,
	}

	s.Services = *NewServiceFacade(zerolog.New(os.Stdout), repoFacade)
}

func Test_RunServicesTestSuite(t *testing.T) {
	suite.Run(t, new(ServicesTestSuite))
}
