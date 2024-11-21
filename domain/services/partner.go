package services

import (
	"context"

	"github.com/JamshedJ/WalletAPI/domain/entities"
	"github.com/JamshedJ/WalletAPI/domain/repository"
	"github.com/rs/zerolog"
)

type PartnerServiceI interface {
	GetPartnerByID(ctx context.Context, partnerID string) (*entities.Partner, error)
}

var _ PartnerServiceI = (*PartnerService)(nil)

type PartnerService struct {
	Repo   repository.PartnerRepositoryI
	Logger zerolog.Logger
}

func (s *PartnerService) GetPartnerByID(ctx context.Context, partnerID string) (*entities.Partner, error) {
	logger := s.Logger.With().Str("partner_id", partnerID).Logger()
	partner, err := s.Repo.GetPartnerByID(ctx, partnerID)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to get partner")
		return nil, err
	}
	return partner, nil
}
