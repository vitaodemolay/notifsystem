package campaign

import (
	contract "github.com/vitaodemolay/notifsystem/internal/application/contract/campaign"
)

type CampaignService interface{
	CreateCampaign(request *contract.CreateCampaign) error
}