package gormRepo

import (
	"context"

	"github.com/JamshedJ/WalletAPI/domain/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Partner struct {
	ID        uuid.UUID `gorm:"primaryKey"`
	Name      string
	SecretKey string
}

func (p *Partner) ToEntity() *entities.Partner {
	return &entities.Partner{
		ID:        p.ID,
		Name:      p.Name,
		SecretKey: p.SecretKey,
	}
}

type GormPartnerRepo struct {
	DB *gorm.DB
}

func (r *GormPartnerRepo) GetPartnerByID(ctx context.Context, id string) (*entities.Partner, error) {
	var p Partner
	err := r.DB.WithContext(ctx).First(&p, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return p.ToEntity(), nil
}
