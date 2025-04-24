package campaign

import (
	"errors"

	contract "github.com/vitaodemolay/notifsystem/internal/application/contract/campaign"
	repository "github.com/vitaodemolay/notifsystem/internal/domain/repository/campaign"
)

//go:generate go run go.uber.org/mock/mockgen -package=mock -destination=./mock/campaign.go . CampaignService
type CampaignService interface {
	CreateCampaign(request *contract.CreateCampaign) (string, error)
}

type campaingService struct {
	campaignRepository repository.CampaignRepository
}

func NewCampaignService(campaignRepository repository.CampaignRepository) (CampaignService, error) {
	if campaignRepository == nil {
		return nil, errors.New("campaign repository is nil")
	}

	return &campaingService{
		campaignRepository: campaignRepository,
	}, nil
}

func (c *campaingService) CreateCampaign(request *contract.CreateCampaign) (string, error) {
	if campaign, err := MapToDomain(request); err != nil {
		return "", err
	} else if err = c.campaignRepository.Save(campaign); err != nil {
		return "", err
	} else {
		return campaign.ID, nil
	}
}
