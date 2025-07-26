package container

import (
	"github.com/vitaodemolay/notifsystem/internal/infrastructure/web/controllers/basictest"
	"github.com/vitaodemolay/notifsystem/internal/infrastructure/web/controllers/campaign"
	"github.com/vitaodemolay/notifsystem/internal/infrastructure/web/entrypoint"
)

type EntryPoint struct {
	Campaign  *campaign.Controller
	BasicTest *basictest.Controller
}

func NewEntryPointContainer(appContainer *ApplicationContainer) (*EntryPoint, error) {
	return &EntryPoint{
		BasicTest: basictest.NewController(),
		Campaign:  campaign.NewController(appContainer.CampaignService),
	}, nil
}

func (e *EntryPoint) GetControllers() []entrypoint.Router {
	return []entrypoint.Router{
		e.BasicTest,
		e.Campaign,
	}
}
