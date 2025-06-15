package repository

import (
	"github.com/vitaodemolay/notifsystem/internal/domain/model/campaign"
	"github.com/vitaodemolay/notifsystem/internal/infrastructure/database/dbcontext"
	internalerrors "github.com/vitaodemolay/notifsystem/pkg/internal-errors"
	"gorm.io/gorm"
)

type campaignRepository struct {
	database *gorm.DB
}

func NewCampaignRepository(connectionString string) (*campaignRepository, error) {
	db, err := dbcontext.NewPgDb(connectionString)
	if err != nil {
		return nil, err
	}
	return &campaignRepository{
		database: db,
	}, nil
}

func (r *campaignRepository) Create(campaign *campaign.Campaign) error {
	return r.database.Create(campaign).Error
}

func (r *campaignRepository) FindByID(id string) (*campaign.Campaign, error) {
	var campaign campaign.Campaign
	if err := r.database.
		Preload("Contacts").
		Preload("StatusList").
		First(&campaign, "id = ?", id).Error; err != nil && err == gorm.ErrRecordNotFound {
		return nil, internalerrors.ErrNotFound
	} else if err != nil {
		return nil, err
	}
	return &campaign, nil
}

func (r *campaignRepository) FindAll() ([]campaign.Campaign, error) {
	var campaigns []campaign.Campaign
	if err := r.database.
		Preload("Contacts").
		Preload("StatusList").
		Find(&campaigns).Error; err != nil {
		return nil, err
	}
	return campaigns, nil
}
