package campaign

import (
	"errors"
	"testing"
	"time"

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
	// Arrange
	assert := assert.New(t)

	// Act
	campaignService, err := NewCampaignService(nil)

	// Assert
	assert.Nil(campaignService)
	assert.Equal(err.Error(), "campaign repository is nil")
}

func Test_NewCampaignService_WhenCampaignRepositoryIsNotNil(t *testing.T) {
	// Arrange
	assert := assert.New(t)
	suite := setup(t)

	// Act
	campaignService, err := NewCampaignService(suite.Repo)

	// Assert
	assert.NotNil(campaignService)
	assert.Nil(err)
}

func Test_CreateCampaign_WhenRequestIsNil(t *testing.T) {
	// Arrange
	assert := assert.New(t)
	suite := setup(t)
	campaignService, _ := NewCampaignService(suite.Repo)

	// Act
	campaignID, err := campaignService.CreateCampaign(nil)

	// Assert
	assert.Equal(campaignID, "")
	assert.Equal(err.Error(), "request is nil")
}

func Test_CreateCampaign_WhenRequestIsValid(t *testing.T) {
	// Arrange
	assert := assert.New(t)
	suite := setup(t)
	request := suite.Request
	campaignService, _ := NewCampaignService(suite.Repo)

	suite.Repo.EXPECT().Create(gomock.Cond(func(c *campaign.Campaign) bool {
		return c.Title == request.Title &&
			c.Content == request.Content &&
			len(c.Contacts) == len(request.Emails) &&
			c.ID != ""
	})).Return(nil).Times(1)

	// Act
	campaignID, err := campaignService.CreateCampaign(request)

	// Assert
	assert.NotEqual(campaignID, "")
	assert.Nil(err)
}

func Test_CreateCampaign_WhenRequestIsInvalid(t *testing.T) {
	// Arrange
	assert := assert.New(t)
	suite := setup(t)
	request := suite.Request
	campaignService, _ := NewCampaignService(suite.Repo)

	request.Emails = []string{}

	// Act
	campaignID, err := campaignService.CreateCampaign(request)

	// Assert
	assert.Equal(campaignID, "")
	assert.Equal(err.Error(), "contacts is required with min 1")
}

func Test_CreateCampaign_WhenRepositoryFails(t *testing.T) {
	// Arrange
	assert := assert.New(t)
	suite := setup(t)
	request := suite.Request
	campaignService, _ := NewCampaignService(suite.Repo)
	testErrorMessage := "error to test"

	suite.Repo.EXPECT().Create(gomock.Any()).Return(errors.New(testErrorMessage)).Times(1)

	// Act
	campaignID, err := campaignService.CreateCampaign(request)

	// Assert
	assert.Equal(campaignID, "")
	assert.NotEqual(err.Error(), testErrorMessage)
	assert.Equal(err.Error(), internalerrors.ErrInternal.Error())
}

func Test_GetCampaigns_WhenRepositoryFails(t *testing.T) {
	// Arrange
	assert := assert.New(t)
	suite := setup(t)
	campaignService, _ := NewCampaignService(suite.Repo)
	testErrorMessage := "error to test"

	suite.Repo.EXPECT().FindAll().Return(nil, errors.New(testErrorMessage)).Times(1)

	// Act
	campaigns, err := campaignService.GetCampaigns()

	// Assert
	assert.Nil(campaigns)
	assert.NotEqual(err.Error(), testErrorMessage)
	assert.ErrorIs(err, internalerrors.ErrInternal)
}

func Test_GetCampaigns_WhenRepositoryReturnsEmptyList(t *testing.T) {
	// Arrange
	assert := assert.New(t)
	suite := setup(t)
	campaignService, _ := NewCampaignService(suite.Repo)

	suite.Repo.EXPECT().FindAll().Return([]campaign.Campaign{}, nil).Times(1)

	// Act
	campaigns, err := campaignService.GetCampaigns()

	// Assert
	assert.Nil(err)
	assert.NotNil(campaigns)
	assert.Empty(campaigns, "Expected campaigns to be empty")
}

func Test_GetCampaigns_WhenRepositoryReturnsList(t *testing.T) {
	// Arrange
	assert := assert.New(t)
	suite := setup(t)
	campaignService, _ := NewCampaignService(suite.Repo)

	expectedCampaigns := []campaign.Campaign{
		{
			ID:        "1",
			Title:     "Campaign 1",
			CreatedAt: time.Now(),
			Content:   "Content 1",
			Contacts:  []campaign.Contact{{Email: "email1@test.com"}},
		},
		{
			ID:        "2",
			Title:     "Campaign 2",
			CreatedAt: time.Now(),
			Content:   "Content 2",
			Contacts:  []campaign.Contact{{Email: "email2@test.com"}},
		},
	}

	suite.Repo.EXPECT().FindAll().Return(expectedCampaigns, nil).Times(1)

	// Act
	campaigns, err := campaignService.GetCampaigns()

	// Assert
	assert.Nil(err)
	assert.NotNil(campaigns)
	assert.Len(campaigns, len(expectedCampaigns), "Expected campaigns to match the expected list")
	for i, campaign := range campaigns {
		assert.Equal(expectedCampaigns[i].ID, campaign.ID, "Expected campaign ID to match")
		assert.Equal(expectedCampaigns[i].Title, campaign.Title, "Expected campaign title to match")
		assert.Equal(expectedCampaigns[i].Content, campaign.Content, "Expected campaign content to match")
		assert.Len(campaign.Emails, len(expectedCampaigns[i].Contacts), "Expected contacts length to match")
		for j, contact := range campaign.Emails {
			assert.Equal(expectedCampaigns[i].Contacts[j].Email, contact, "Expected contact email to match")
		}
	}
}

func Test_GetCampaignByID_WhenIDIsEmpty(t *testing.T) {
	// Arrange
	assert := assert.New(t)
	suite := setup(t)
	campaignService, _ := NewCampaignService(suite.Repo)

	// Act
	campaign, err := campaignService.GetCampaignByID("")

	// Assert
	assert.Nil(campaign)
	assert.ErrorIs(err, internalerrors.ErrBadRequest)
}

func Test_GetCampaignByID_WhenRepositoryReturnsError(t *testing.T) {
	// Arrange
	assert := assert.New(t)
	suite := setup(t)
	campaignService, _ := NewCampaignService(suite.Repo)
	testErrorMessage := "error to test"

	suite.Repo.EXPECT().FindByID("1").Return(nil, errors.New(testErrorMessage)).Times(1)

	// Act
	campaign, err := campaignService.GetCampaignByID("1")

	// Assert
	assert.Nil(campaign)
	assert.NotEqual(err.Error(), testErrorMessage)
	assert.ErrorIs(err, internalerrors.ErrInternal)
}

func Test_GetCampaignByID_WhenRepositoryReturnsNotFound(t *testing.T) {
	// Arrange
	assert := assert.New(t)
	suite := setup(t)
	campaignService, _ := NewCampaignService(suite.Repo)

	suite.Repo.EXPECT().FindByID("1").Return(nil, internalerrors.ErrNotFound).Times(1)

	// Act
	campaign, err := campaignService.GetCampaignByID("1")

	// Assert
	assert.Nil(campaign)
	assert.ErrorIs(err, internalerrors.ErrNotFound)
}

func Test_GetCampaignByID_WhenRepositoryReturnsCampaign(t *testing.T) {
	// Arrange
	assert := assert.New(t)
	suite := setup(t)
	campaignService, _ := NewCampaignService(suite.Repo)

	expectedCampaign := &campaign.Campaign{
		ID:        "1",
		Title:     "Campaign 1",
		CreatedAt: time.Now(),
		Content:   "Content 1",
		Contacts:  []campaign.Contact{{Email: "email1@test.com"}},
	}
	suite.Repo.EXPECT().FindByID("1").Return(expectedCampaign, nil).Times(1)

	// Act
	campaign, err := campaignService.GetCampaignByID("1")

	// Assert
	assert.Nil(err)
	assert.NotNil(campaign)
	assert.Equal(expectedCampaign.ID, campaign.ID, "Expected campaign ID to match")
	assert.Equal(expectedCampaign.Title, campaign.Title, "Expected campaign title to match")
	assert.Equal(expectedCampaign.Content, campaign.Content, "Expected campaign content to match")
	assert.Len(campaign.Emails, len(expectedCampaign.Contacts), "Expected contacts length to match")
	assert.Equal(expectedCampaign.CreatedAt.Format(time.RFC3339), campaign.CreatedAt, "Expected campaign created at to match")
	for i, contact := range campaign.Emails {
		assert.Equal(expectedCampaign.Contacts[i].Email, contact, "Expected contact email to match")
	}
}
