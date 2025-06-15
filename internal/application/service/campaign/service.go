package campaign

import (
	"errors"

	contract "github.com/vitaodemolay/notifsystem/internal/application/contract/campaign"
	repository "github.com/vitaodemolay/notifsystem/internal/domain/repository/campaign"
	internalerrors "github.com/vitaodemolay/notifsystem/pkg/internal-errors"
)

//go:generate go run go.uber.org/mock/mockgen -package=mock -destination=./mock/campaign.go . CampaignService
type CampaignService interface {
	CreateCampaign(request *contract.CreateCampaign) (string, error)
	GetCampaigns() ([]*contract.Campaign, error)
	GetCampaignByID(id string) (*contract.Campaign, error)
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
	} else if err = c.campaignRepository.Create(campaign); err != nil {
		return "", internalerrors.ErrInternal
	} else {
		return campaign.ID, nil
	}
}

func (c *campaingService) GetCampaigns() ([]*contract.Campaign, error) {
	campaigns, err := c.campaignRepository.FindAll()
	if err != nil {
		return nil, internalerrors.ErrInternal
	}

	return MapToContractList(campaigns), nil
}

func (c *campaingService) GetCampaignByID(id string) (*contract.Campaign, error) {
	if id == "" {
		return nil, internalerrors.ErrBadRequest
	}

	campaign, err := c.campaignRepository.FindByID(id)
	if err != nil {
		if errors.Is(err, internalerrors.ErrNotFound) {
			return nil, err
		}
		return nil, internalerrors.ErrInternal
	}

	return MapToContract(campaign), nil
}
