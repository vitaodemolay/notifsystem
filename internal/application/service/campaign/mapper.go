package campaign

import (
	"errors"

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
