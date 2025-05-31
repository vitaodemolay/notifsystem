package container

import (
	domain "github.com/vitaodemolay/notifsystem/internal/domain/model/campaign"
	repository "github.com/vitaodemolay/notifsystem/internal/domain/repository/campaign"
	internalerrors "github.com/vitaodemolay/notifsystem/pkg/internal-errors"
)

type InfraContainer struct {
	CampaignRepository repository.CampaignRepository
}

func NewInfraContainer() (*InfraContainer, error) {
	campaignRepo := &fakeCampaignRepository{
		repo: make(map[string]domain.Campaign),
	}

	return &InfraContainer{
		CampaignRepository: campaignRepo,
	}, nil
}

type fakeCampaignRepository struct {
	repo map[string]domain.Campaign
}

func (f *fakeCampaignRepository) Save(campaign *domain.Campaign) error {
	f.repo[campaign.ID] = *campaign
	return nil
}

func (f *fakeCampaignRepository) FindByID(id string) (*domain.Campaign, error) {
	campaign, exists := f.repo[id]
	if !exists {
		return nil, internalerrors.ErrNotFound
	}
	return &campaign, nil
}

func (f *fakeCampaignRepository) FindAll() ([]domain.Campaign, error) {
	var campaigns []domain.Campaign
	for _, campaign := range f.repo {
		campaigns = append(campaigns, campaign)
	}
	return campaigns, nil
}
