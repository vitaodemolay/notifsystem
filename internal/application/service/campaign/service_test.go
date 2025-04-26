package campaign

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	contract "github.com/vitaodemolay/notifsystem/internal/application/contract/campaign"
	"github.com/vitaodemolay/notifsystem/internal/domain/model/campaign"
	"github.com/vitaodemolay/notifsystem/internal/domain/repository/campaign/mock"
	internalerrors "github.com/vitaodemolay/notifsystem/pkg/internal-errors"
	"go.uber.org/mock/gomock"
)

type testSuitService struct {
	Request *contract.CreateCampaign
	Control *gomock.Controller
	Repo    *mock.MockCampaignRepository
}

func setup(t *testing.T) *testSuitService {
	ctrl := gomock.NewController(t)
	repo := mock.NewMockCampaignRepository(ctrl)

	return &testSuitService{
		Request: &contract.CreateCampaign{
			Title:   "Campaign X of Test",
			Content: "Body of Campaign X of Test",
			Emails:  []string{"email1@test.com", "email2@test.com"},
		},
		Control: ctrl,
		Repo:    repo,
	}
}

func Test_NewCampaignService_WhenCampaignRepositoryIsNil(t *testing.T) {
	assert := assert.New(t)

	campaignService, err := NewCampaignService(nil)

	assert.Nil(campaignService)
	assert.Equal(err.Error(), "campaign repository is nil")
}

func Test_NewCampaignService_WhenCampaignRepositoryIsNotNil(t *testing.T) {
	assert := assert.New(t)
	suite := setup(t)

	campaignService, err := NewCampaignService(suite.Repo)

	assert.NotNil(campaignService)
	assert.Nil(err)
}

func Test_CreateCampaign_WhenRequestIsNil(t *testing.T) {
	assert := assert.New(t)
	suite := setup(t)
	campaignService, _ := NewCampaignService(suite.Repo)

	campaignID, err := campaignService.CreateCampaign(nil)

	assert.Equal(campaignID, "")
	assert.Equal(err.Error(), "request is nil")
}

func Test_CreateCampaign_WhenRequestIsValid(t *testing.T) {
	assert := assert.New(t)
	suite := setup(t)
	request := suite.Request
	campaignService, _ := NewCampaignService(suite.Repo)

	suite.Repo.EXPECT().Save(gomock.Cond(func(c *campaign.Campaign) bool {
		return c.Title == request.Title &&
			c.Content == request.Content &&
			len(c.Contacts) == len(request.Emails) &&
			c.ID != ""
	})).Return(nil).Times(1)

	campaignID, err := campaignService.CreateCampaign(request)

	assert.NotEqual(campaignID, "")
	assert.Nil(err)
}

func Test_CreateCampaign_WhenRequestIsInvalid(t *testing.T) {
	assert := assert.New(t)
	suite := setup(t)
	request := suite.Request
	campaignService, _ := NewCampaignService(suite.Repo)

	request.Emails = []string{}

	campaignID, err := campaignService.CreateCampaign(request)

	assert.Equal(campaignID, "")
	assert.Equal(err.Error(), "contacts is required with min 1")
}

func Test_CreateCampaign_WhenRepositoryFails(t *testing.T) {
	assert := assert.New(t)
	suite := setup(t)
	request := suite.Request
	campaignService, _ := NewCampaignService(suite.Repo)
	testErrorMessage := "error to test"

	suite.Repo.EXPECT().Save(gomock.Any()).Return(errors.New(testErrorMessage)).Times(1)

	campaignID, err := campaignService.CreateCampaign(request)

	assert.Equal(campaignID, "")
	assert.NotEqual(err.Error(), testErrorMessage)
	assert.Equal(err.Error(), internalerrors.ErrInternal.Error())
}
