package campaign

import "github.com/vitaodemolay/notifsystem/internal/domain/model/campaign"

//go:generate go run go.uber.org/mock/mockgen -package=mock -destination=./mock/campaign.go . CampaignRepository
type CampaignRepository interface {
	Save(campaign *campaign.Campaign) error
	FindByID(id string) (*campaign.Campaign, error)
	FindAll() ([]campaign.Campaign, error)
}
