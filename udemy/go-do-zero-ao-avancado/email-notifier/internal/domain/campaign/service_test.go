package campaign_test

import (
	"errors"
	"testing"

	"github.com/igorsalves/estudos/tree/main/udemy/go-do-zero-ao-avancado/email-notifier/internal/contract"
	"github.com/igorsalves/estudos/tree/main/udemy/go-do-zero-ao-avancado/email-notifier/internal/domain/campaign"
	internalerrors "github.com/igorsalves/estudos/tree/main/udemy/go-do-zero-ao-avancado/email-notifier/internal/internal-errors"
	internalmock "github.com/igorsalves/estudos/tree/main/udemy/go-do-zero-ao-avancado/email-notifier/internal/test/internal-mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	newCampaign = contract.NewCampaign{
		Name:    "Test Y",
		Content: "Body Hi!",
		Emails:  []string{"teste1@test.com"},
	}
	service = campaign.ServiceImp{}
)

func Test_Create_Campaign(t *testing.T) {
	assert := assert.New(t)
	repositoryMock := new(internalmock.CampaignRepositoryMock)
	repositoryMock.On("Create", mock.Anything).Return(nil)
	service.Repository = repositoryMock

	id, err := service.Create(newCampaign)

	assert.NotNil(id)
	assert.Nil(err)
}

func Test_Create_ValidateDomainError(t *testing.T) {
	assert := assert.New(t)

	_, err := service.Create(contract.NewCampaign{})

	assert.False(errors.Is(internalerrors.ErrInternal, err))
}

func Test_Create_SaveCampaign(t *testing.T) {
	repositoryMock := new(internalmock.CampaignRepositoryMock)
	repositoryMock.On("Create", mock.MatchedBy(func(campaign *campaign.Campaign) bool {
		if campaign.Name != newCampaign.Name ||
			campaign.Content != newCampaign.Content ||
			len(campaign.Contacts) != len(newCampaign.Emails) {
			return false
		}

		return true
	})).Return(nil)
	service.Repository = repositoryMock

	service.Create(newCampaign)

	repositoryMock.AssertExpectations(t)
}

func Test_Create_ValidateRepositorySave(t *testing.T) {
	assert := assert.New(t)
	repositoryMock := new(internalmock.CampaignRepositoryMock)
	repositoryMock.On("Create", mock.Anything).Return(errors.New("error to save on database"))
	service.Repository = repositoryMock

	_, err := service.Create(newCampaign)

	assert.True(errors.Is(internalerrors.ErrInternal, err))
}

func Test_GetById_ReturnCampaign(t *testing.T) {
	assert := assert.New(t)
	repositoryMock := new(internalmock.CampaignRepositoryMock)
	campaignExpected, _ := campaign.NewCampaign(
		newCampaign.Name,
		newCampaign.Content,
		newCampaign.Emails,
	)
	repositoryMock.On("GetBy", mock.MatchedBy(func(id string) bool {
		return id == campaignExpected.ID
	})).Return(campaignExpected, nil)
	service.Repository = repositoryMock

	campaignReturned, err := service.GetBy(campaignExpected.ID)

	assert.Nil(err)
	assert.Equal(campaignExpected.ID, campaignReturned.ID)
	assert.Equal(campaignExpected.Name, campaignReturned.Name)
	assert.Equal(campaignExpected.Content, campaignReturned.Content)
	assert.Equal(campaignExpected.Status, campaignReturned.Status)
}

func Test_GetById_ReturnErrorWhenSomethingWrong(t *testing.T) {
	assert := assert.New(t)
	repositoryMock := new(internalmock.CampaignRepositoryMock)

	repositoryMock.On("GetBy", mock.Anything).Return(nil, errors.New("Something Error"))
	service.Repository = repositoryMock

	campaignReturned, err := service.GetBy("some-id")

	assert.True(errors.Is(internalerrors.ErrInternal, err))
	assert.Nil(campaignReturned)
}
