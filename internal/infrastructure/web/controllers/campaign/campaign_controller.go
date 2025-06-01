package campaign

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	contract "github.com/vitaodemolay/notifsystem/internal/application/contract/campaign"
	"github.com/vitaodemolay/notifsystem/internal/application/service/campaign"
	"github.com/vitaodemolay/notifsystem/internal/infrastructure/web/entrypoint"
	internalerrors "github.com/vitaodemolay/notifsystem/pkg/internal-errors"
)

type Controller struct {
	service campaign.CampaignService
}

func NewController(service campaign.CampaignService) *Controller {
	return &Controller{
		service: service,
	}
}

func (c *Controller) Path() string {
	return "/v1/campaign"
}

func (c *Controller) GetRoutes() []entrypoint.Route {
	return []entrypoint.Route{
		{
			Method:  http.MethodGet,
			Pattern: "/{id}",
			Handler: c.GetCampaignByID,
		},
		{
			Method:  http.MethodGet,
			Pattern: "/",
			Handler: c.GetCampaigns,
		},
		{
			Method:  http.MethodPost,
			Pattern: "/",
			Handler: c.CreateCampaign,
		},
	}
}

func (c *Controller) GetCampaignByID(w http.ResponseWriter, r *http.Request) (any, int, error) {
	campaignID := chi.URLParam(r, "id")
	if campaignID == "" {
		return nil, 0, internalerrors.ErrBadRequest
	}

	campaign, err := c.service.GetCampaignByID(campaignID)
	if err != nil {
		return nil, 0, err
	}

	return campaign, http.StatusOK, nil
}

func (c *Controller) GetCampaigns(w http.ResponseWriter, r *http.Request) (any, int, error) {
	campaigns, err := c.service.GetCampaigns()
	if err != nil {
		return nil, 0, err
	}
	return campaigns, http.StatusOK, nil
}

func (c *Controller) CreateCampaign(w http.ResponseWriter, r *http.Request) (any, int, error) {
	var request contract.CreateCampaign
	render.DecodeJSON(r.Body, &request)
	id, err := c.service.CreateCampaign(&request)
	if err != nil {
		return nil, 0, err
	}
	return map[string]string{"campaign_id": id}, http.StatusCreated, err
}
