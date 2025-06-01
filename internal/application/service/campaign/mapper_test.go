package campaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	contract "github.com/vitaodemolay/notifsystem/internal/application/contract/campaign"
	model "github.com/vitaodemolay/notifsystem/internal/domain/model/campaign"
	"github.com/vitaodemolay/notifsystem/internal/domain/model/contact"
)

func fakeCreateCampaing() *contract.CreateCampaign {
	return &contract.CreateCampaign{
		Title:   "Campaign X of Test",
		Content: "Body of Campaign X of Test",
		Emails:  []string{"email1@test.com", "email2@test.com"},
	}
}

func fakeDomainCampaign() *model.Campaign {
	create, _ := time.Parse(time.RFC3339, "2023-10-01T00:00:00Z")
	campaign := &model.Campaign{
		ID:        "123",
		Title:     "Campaign X of Test",
		Content:   "Body of Campaign X of Test",
		CreatedAt: create,
		Contacts:  []contact.Contact{{Email: "teste1@test.com"}},
	}
	return campaign
}

func fakeListDomainCampaign(count int) []model.Campaign {
	campaigns := make([]model.Campaign, count)
	for i := 0; i < count; i++ {
		campaigns[i] = *fakeDomainCampaign()
	}

	return campaigns
}

func Test_MapToDomain_WhenRequestIsNil(t *testing.T) {
	// Arrange
	assert := assert.New(t)

	// Act
	campaign, err := MapToDomain(nil)

	// Assert
	assert.Nil(campaign)
	assert.Equal(err.Error(), "request is nil")
}

func Test_MapToDomain_WhenRequestIsValid(t *testing.T) {
	// Arrange
	assert := assert.New(t)
	request := fakeCreateCampaing()

	// Act
	campaign, err := MapToDomain(request)

	// Assert
	assert.Nil(err)
	assert.Equal(campaign.Title, request.Title)
	assert.Equal(campaign.Content, request.Content)
	assert.Equal(len(campaign.Contacts), len(request.Emails))
}

func Test_MapToDomain_WhenRequestIsInvalid(t *testing.T) {
	// Arrange
	assert := assert.New(t)
	request := fakeCreateCampaing()

	request.Emails = []string{}

	// Act
	campaign, err := MapToDomain(request)

	// Assert
	assert.Nil(campaign)
	assert.Equal(err.Error(), "contacts is required with min 1")
}

func Test_MapToContract_WhenCampaignIsNil(t *testing.T) {
	// Arrange
	assert := assert.New(t)

	// Act
	campaign := MapToContract(nil)

	// Assert
	assert.Nil(campaign)
}

func Test_MapToContract_WhenCampaignIsValid(t *testing.T) {
	// Arrange
	assert := assert.New(t)
	campaign := fakeDomainCampaign()

	// Act
	contract := MapToContract(campaign)

	// Assert
	assert.NotNil(contract)
	assert.Equal(contract.ID, campaign.ID)
	assert.Equal(contract.Title, campaign.Title)
	assert.Equal(contract.Content, campaign.Content)
	assert.Equal(contract.CreatedAt, campaign.CreatedAt.Format(time.RFC3339))
	assert.Equal(len(contract.Emails), len(campaign.Contacts))
}

func Test_MapToContractList_WhenCampaignsIsNil(t *testing.T) {
	// Arrange
	assert := assert.New(t)

	// Act
	contracts := MapToContractList(nil)

	// Assert
	assert.Nil(contracts)
}

func Test_MapToContractList_WhenCampaignsIsValid(t *testing.T) {
	// Arrange
	assert := assert.New(t)
	campaigns := fakeListDomainCampaign(3)

	// Act
	contracts := MapToContractList(campaigns)

	// Assert
	assert.NotNil(contracts)
	assert.Len(contracts, len(campaigns))
	for i, contract := range contracts {
		assert.Equal(contract.ID, campaigns[i].ID)
		assert.Equal(contract.Title, campaigns[i].Title)
		assert.Equal(contract.Content, campaigns[i].Content)
		assert.Equal(contract.CreatedAt, campaigns[i].CreatedAt.Format(time.RFC3339))
		assert.Equal(len(contract.Emails), len(campaigns[i].Contacts))
	}
}
