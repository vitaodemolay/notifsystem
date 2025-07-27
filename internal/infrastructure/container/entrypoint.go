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

func NewEntryPointContainer(appContainer *ApplicationContainer, clientId, redirectUri, tokenType string) (*EntryPoint, error) {
	idProvider := entrypoint.NewIdentityProvider(clientId, redirectUri, tokenType)

	return &EntryPoint{
		BasicTest: basictest.NewController(),
		Campaign:  campaign.NewController(appContainer.CampaignService, idProvider),
	}, nil
}

func (e *EntryPoint) GetControllers() []entrypoint.Router {
	return []entrypoint.Router{
		e.BasicTest,
		e.Campaign,
	}
}
