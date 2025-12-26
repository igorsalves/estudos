package campaign

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCampaign(t *testing.T) {
	name := "Campaign x"
	content := "This is the content of Campaign x."
	contacts := []string{"contact1@example.com", "contact2@example.com"}

	campaign := NewCampaign(name, content, contacts)

	assert := assert.New(t)
	assert.Equal(campaign.ID, "1")
	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(contacts))
}
