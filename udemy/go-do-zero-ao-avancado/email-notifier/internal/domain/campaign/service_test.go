package campaign

import (
	"errors"
	"testing"

	"github.com/igorsalves/estudos/tree/main/udemy/go-do-zero-ao-avancado/email-notifier/internal/contract"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

var (
	newCampaing = contract.NewCampaign{
		Name:    "Black Friday",
		Content: "Don't miss our Black Friday deals.",
		Emails:  []string{"customer1@example.com", "customer2@example.com"},
	}
	repository = new(repositoryMock)
	service    = Service{
		Repository: repository,
	}
)

func Test_Create_Campaign(t *testing.T) {
	assert := assert.New(t)
	repository.On("Save", mock.Anything).Return(nil)

	id, err := service.Create(newCampaing)

	assert.NotEmpty(id)
	assert.Nil(err)
}

func Test_Create_ValidateDomainError(t *testing.T) {
	assert := assert.New(t)
	newCampaing.Name = ""

	_, err := service.Create(newCampaing)

	assert.NotNil(err)
	assert.Equal("name is required", err.Error())
}

func Test_Create_SaveCampaign(t *testing.T) {

	repository.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {
		if campaign.Name != newCampaing.Name {
			return false
		}
		if campaign.Content != newCampaing.Content {
			return false
		}
		if len(campaign.Contacts) != len(newCampaing.Emails) {
			return false
		}
		return true
	})).Return(nil)

	service.Create(newCampaing)

	repository.AssertExpectations(t)
}

func Test_Create_ValidateRepositorySave(t *testing.T) {
	assert := assert.New(t)
	repository.On("Save", mock.Anything).Return(errors.New("database error"))

	_, err := service.Create(newCampaing)

	assert.Equal("database error", err.Error())
}
