package container

import (
	repository "github.com/vitaodemolay/notifsystem/internal/domain/repository/campaign"
	dbcontext "github.com/vitaodemolay/notifsystem/internal/infrastructure/database/repository"
)

type InfraContainer struct {
	CampaignRepository repository.CampaignRepository
}

func NewInfraContainer(connectionString string) (*InfraContainer, error) {
	campaignRepo, err := dbcontext.NewCampaignRepository(connectionString)
	if err != nil {
		return nil, err
	}

	return &InfraContainer{
		CampaignRepository: campaignRepo,
	}, nil
}

// func newfakeCampaignRepository() repository.CampaignRepository {
// 	return &fakeCampaignRepository{
// 		repo: make(map[string]domain.Campaign),
// 	}
// }

// type fakeCampaignRepository struct {
// 	repo map[string]domain.Campaign
// }

// func (f *fakeCampaignRepository) Create(campaign *domain.Campaign) error {
// 	f.repo[campaign.ID] = *campaign
// 	return nil
// }

// func (f *fakeCampaignRepository) FindByID(id string) (*domain.Campaign, error) {
// 	campaign, exists := f.repo[id]
// 	if !exists {
// 		return nil, internalerrors.ErrNotFound
// 	}
// 	return &campaign, nil
// }

// func (f *fakeCampaignRepository) FindAll() ([]domain.Campaign, error) {
// 	var campaigns []domain.Campaign
// 	for _, campaign := range f.repo {
// 		campaigns = append(campaigns, campaign)
// 	}
// 	return campaigns, nil
// }
