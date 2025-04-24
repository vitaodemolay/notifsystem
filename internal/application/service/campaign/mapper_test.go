package campaign

import (
	"testing"

	"github.com/stretchr/testify/assert"
	contract "github.com/vitaodemolay/notifsystem/internal/application/contract/campaign"
)

type testSuitMapper struct {
	Request *contract.CreateCampaign
}

func setuptestSuitMapper(t *testing.T) *testSuitMapper {

	return &testSuitMapper{
		Request: &contract.CreateCampaign{
			Title:   "title",
			Content: "content",
			Emails:  []string{"email1@test.com", "email2@test.com"},
		},
	}
}

func Test_MapToDomain_WhenRequestIsNil(t *testing.T) {
	assert := assert.New(t)

	campaign, err := MapToDomain(nil)

	assert.Nil(campaign)
	assert.Equal(err.Error(), "request is nil")
}

func Test_MapToDomain_WhenRequestIsValid(t *testing.T) {
	assert := assert.New(t)
	suite := setuptestSuitMapper(t)
	request := suite.Request

	campaign, err := MapToDomain(request)

	assert.Nil(err)
	assert.Equal(campaign.Title, request.Title)
	assert.Equal(campaign.Content, request.Content)
	assert.Equal(len(campaign.Contacts), len(request.Emails))
}

func Test_MapToDomain_WhenRequestIsInvalid(t *testing.T) {
	assert := assert.New(t)
	suite := setuptestSuitMapper(t)
	request := suite.Request

	request.Emails = []string{}

	campaign, err := MapToDomain(request)

	assert.Nil(campaign)
	assert.Equal(err.Error(), "at least one email is required")
}
