package campaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	title    = "Campaign X"
	content  = "Body"
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
	assert.Equal(err.Error(), "title cannot be empty")
}

func Test_NewCampaign_EmptyContent(t *testing.T) {
	assert := assert.New(t)

	campaign, err := NewCampaign(title, "", contacts)

	assert.Nil(campaign)
	assert.Equal(err.Error(), "content cannot be empty")
}

func Test_NewCampaign_EmptyContacts(t *testing.T) {
	assert := assert.New(t)

	campaign, err := NewCampaign(title, content, []string{})

	assert.Nil(campaign)
	assert.Equal(err.Error(), "at least one email is required")
}
