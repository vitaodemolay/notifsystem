package campaign

import (
	"testing"
	"time"

	"github.com/jaswdr/faker/v2"
	"github.com/stretchr/testify/assert"
)

var (
	title    = "Campaign X of Test"
	content  = "Body of Campaign X of Test"
	contacts = []string{"email1@e.com", "email2@e.com"}
)

func Test_NewCampaign_CreateCampaign(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaign(title, content, contacts)

	assert.Equal(campaign.Title, title)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(contacts))
}

func Test_NewCampaign_IDIsNotNill(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaign(title, content, contacts)

	assert.NotNil(campaign.ID)
}

func Test_NewCampaign_CreateAtMustBeNow(t *testing.T) {
	assert := assert.New(t)
	now := time.Now().Add(-time.Minute)

	campaign, _ := NewCampaign(title, content, contacts)

	assert.Greater(campaign.CreatedAt, now)
}

func Test_NewCampaign_EmptyTitle(t *testing.T) {
	assert := assert.New(t)

	campaign, err := NewCampaign("", content, contacts)

	assert.Nil(campaign)
	assert.Equal(err.Error(), "title is required with min 10")
}

func Test_NewCampaign_TitleTooLong(t *testing.T) {
	assert := assert.New(t)
	faker := faker.New()

	campaign, err := NewCampaign(faker.Lorem().Text(50), content, contacts)

	assert.Nil(campaign)
	assert.Equal(err.Error(), "title is required with max 30")
}

func Test_NewCampaign_TitleTooShort(t *testing.T) {
	assert := assert.New(t)
	faker := faker.New()

	campaign, err := NewCampaign(faker.Lorem().Text(5), content, contacts)

	assert.Nil(campaign)
	assert.Equal(err.Error(), "title is required with min 10")
}

func Test_NewCampaign_EmptyContent(t *testing.T) {
	assert := assert.New(t)

	campaign, err := NewCampaign(title, "", contacts)

	assert.Nil(campaign)
	assert.Equal(err.Error(), "content is required with min 10")
}

func Test_NewCampaign_ContentTooLong(t *testing.T) {
	assert := assert.New(t)
	faker := faker.New()

	campaign, err := NewCampaign(title, faker.Lorem().Sentence(2000), contacts)

	assert.Nil(campaign)
	assert.Equal(err.Error(), "content is required with max 2048")
}

func Test_NewCampaign_ContentTooShort(t *testing.T) {
	assert := assert.New(t)
	faker := faker.New()

	campaign, err := NewCampaign(title, faker.Lorem().Text(5), contacts)

	assert.Nil(campaign)
	assert.Equal(err.Error(), "content is required with min 10")
}

func Test_NewCampaign_EmptyContacts(t *testing.T) {
	assert := assert.New(t)

	campaign, err := NewCampaign(title, content, []string{})

	assert.Nil(campaign)
	assert.Equal(err.Error(), "contacts is required with min 1")
}

func Test_NewCampaign_EmptyContactEmail(t *testing.T) {
	assert := assert.New(t)

	campaign, err := NewCampaign(title, content, []string{"", ""})
	assert.Nil(campaign)
	assert.Equal(err.Error(), "email cannot be empty")
}

func Test_NewCampaign_ContactEmailInvalid(t *testing.T) {
	assert := assert.New(t)

	campaign, err := NewCampaign(title, content, []string{"email1"})
	assert.Nil(campaign)
	assert.Equal(err.Error(), "email is invalid")
}

func Test_NewCampaign_StatusMustBePending(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaign(title, content, contacts)
	assert.Equal(Pending, campaign.Status)
}
