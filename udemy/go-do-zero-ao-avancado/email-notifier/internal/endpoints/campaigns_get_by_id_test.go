package endpoints

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/igorsalves/estudos/tree/main/udemy/go-do-zero-ao-avancado/email-notifier/internal/contract"
	internalmock "github.com/igorsalves/estudos/tree/main/udemy/go-do-zero-ao-avancado/email-notifier/internal/test/internal-mock"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_CampaignsGetById_should_return_campaign(t *testing.T) {
	assert := assert.New(t)
	campaignResponse := contract.CampaignResponse{
		ID:      "321",
		Name:    "teste",
		Content: "Hi everyone",
		Status:  "Pending",
	}
	service := new(internalmock.CampaignServiceMock)
	service.On("GetBy", mock.Anything).Return(&campaignResponse, nil)
	handler := Handler{CampaignService: service}
	req, _ := http.NewRequest("GET", "/321", nil)
	rr := httptest.NewRecorder()

	response, status, _ := handler.CampaignGetById(rr, req)

	assert.Equal(200, status)
	assert.Equal(campaignResponse.ID, response.(*contract.CampaignResponse).ID)
	assert.Equal(campaignResponse.Name, response.(*contract.CampaignResponse).Name)
}

func Test_CampaignsGetById_should_return_error_when_something_wrong(t *testing.T) {
	assert := assert.New(t)
	service := new(internalmock.CampaignServiceMock)
	errExpected := errors.New("something wrong")
	service.On("GetBy", mock.Anything).Return(nil, errExpected)
	handler := Handler{CampaignService: service}
	req, _ := http.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	_, _, errReturned := handler.CampaignGetById(rr, req)

	assert.Equal(errExpected.Error(), errReturned.Error())

}
