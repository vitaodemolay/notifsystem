package container

import (
	"github.com/vitaodemolay/notifsystem/internal/application/service/campaign"
)

type ApplicationContainer struct {
	CampaignService campaign.CampaignService
}

func NewApplicationContainer(infraContainer *InfraContainer) (*ApplicationContainer, error) {
	campaignService, err := campaign.NewCampaignService(infraContainer.CampaignRepository)
	if err != nil {
		return nil, err
	}

	return &ApplicationContainer{
		CampaignService: campaignService,
	}, nil
}
