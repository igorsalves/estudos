package campaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_NewCampaign_CreateCampaign(t *testing.T) {
	name := "Campaign x"
	content := "This is the content of Campaign x."
	contacts := []string{"contact1@example.com", "contact2@example.com"}

	campaign := NewCampaign(name, content, contacts)

	assert := assert.New(t)
	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(contacts))
}

func Test_NewCampaign_IDIsNotNil(t *testing.T) {
	name := "Campaign x"
	content := "This is the content of Campaign x."
	contacts := []string{"contact1@example.com", "contact2@example.com"}

	campaign := NewCampaign(name, content, contacts)

	assert := assert.New(t)
	assert.NotNil(campaign.ID)
}

func Test_NewCampaign_CreatedOnIsNotNil(t *testing.T) {
	name := "Campaign x"
	content := "This is the content of Campaign x."
	contacts := []string{"contact1@example.com", "contact2@example.com"}
	now := time.Now().Add(-1 * time.Minute)

	campaign := NewCampaign(name, content, contacts)

	assert := assert.New(t)
	assert.NotNil(campaign.CreatedOn)
	assert.True(campaign.CreatedOn.After(now))
}
