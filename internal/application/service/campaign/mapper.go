package campaign

import (
	"errors"
	"time"

	contract "github.com/vitaodemolay/notifsystem/internal/application/contract/campaign"
	model "github.com/vitaodemolay/notifsystem/internal/domain/model/campaign"
)

func MapToDomain(r *contract.CreateCampaign) (*model.Campaign, error) {
	if r == nil {
		return nil, errors.New("request is nil")
	} else if domain, err := model.NewCampaign(r.Title, r.Content, r.Emails); err != nil {
		return nil, err
	} else {
		return domain, nil
	}
}

func MapToContract(campaign *model.Campaign) *contract.Campaign {
	if campaign == nil {
		return nil
	}

	contract := &contract.Campaign{
		ID:        campaign.ID,
		Title:     campaign.Title,
		Content:   campaign.Content,
		CreatedAt: campaign.CreatedAt.Format(time.RFC3339),
		Emails:    make([]string, len(campaign.Contacts)),
	}
	for i, contact := range campaign.Contacts {
		contract.Emails[i] = contact.Email
	}
	return contract
}

func MapToContractList(campaigns []model.Campaign) []*contract.Campaign {
	if campaigns == nil {
		return nil
	}

	contractList := make([]*contract.Campaign, len(campaigns))
	for i, campaign := range campaigns {
		contractList[i] = MapToContract(&campaign)
	}

	return contractList
}
