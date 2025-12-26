package campaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	name     = "Campaign x"
	content  = "This is the content of Campaign x."
	contacts = []string{"contact1@example.com", "contact2@example.com"}
)

func Test_NewCampaign_CreateCampaign(t *testing.T) {
	campaign, _ := NewCampaign(name, content, contacts)

	assert := assert.New(t)
	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(contacts))
}

func Test_NewCampaign_IDIsNotNil(t *testing.T) {
	campaign, _ := NewCampaign(name, content, contacts)

	assert := assert.New(t)
	assert.NotNil(campaign.ID)
}

func Test_NewCampaign_CreatedOnMustBeNow(t *testing.T) {
	now := time.Now().Add(-1 * time.Minute)

	campaign, _ := NewCampaign(name, content, contacts)

	assert := assert.New(t)
	assert.NotNil(campaign.CreatedOn)
	assert.True(campaign.CreatedOn.After(now))
}

func Test_NewCampaign_MustValidateName(t *testing.T) {

	_, err := NewCampaign("", content, contacts)

	assert := assert.New(t)
	assert.Equal("name is required", err.Error())
}

func Test_NewCampaign_MustValidateContent(t *testing.T) {

	_, err := NewCampaign(name, "", contacts)

	assert := assert.New(t)
	assert.Equal("content is required", err.Error())
}

func Test_NewCampaign_MustValidateContacts(t *testing.T) {

	_, err := NewCampaign(name, content, []string{})

	assert := assert.New(t)
	assert.Equal("at least one contact is required", err.Error())
}
