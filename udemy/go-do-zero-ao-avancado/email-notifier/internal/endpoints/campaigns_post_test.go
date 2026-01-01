package endpoints

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/igorsalves/estudos/tree/main/udemy/go-do-zero-ao-avancado/email-notifier/internal/contract"
	internalmock "github.com/igorsalves/estudos/tree/main/udemy/go-do-zero-ao-avancado/email-notifier/internal/test/internal-mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_CampaignsPost_should_save_new_campaign(t *testing.T) {
	assert := assert.New(t)
	body := contract.NewCampaign{
		Name:    "teste",
		Content: "Hi everyone",
		Emails:  []string{"teste@teste.com"},
	}
	service := new(internalmock.CampaignServiceMock)
	service.On("Create", mock.MatchedBy(func(request contract.NewCampaign) bool {
		if request.Name == body.Name && request.Content == body.Content {
			return true
		} else {
			return false
		}
	})).Return("34x", nil)
	handler := Handler{CampaignService: service}

	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(body)
	req, _ := http.NewRequest("POST", "/", &buf)
	rr := httptest.NewRecorder()

	_, status, err := handler.CampaignPost(rr, req)

	assert.Equal(201, status)
	assert.Nil(err)
}

func Test_CampaignsPost_should_inform_error_when_exist(t *testing.T) {
	assert := assert.New(t)
	body := contract.NewCampaign{
		Name:    "teste",
		Content: "Hi everyone",
		Emails:  []string{"teste@teste.com"},
	}
	service := new(internalmock.CampaignServiceMock)
	service.On("Create", mock.Anything).Return("", fmt.Errorf("error"))
	handler := Handler{CampaignService: service}

	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(body)
	req, _ := http.NewRequest("POST", "/", &buf)
	rr := httptest.NewRecorder()

	_, _, err := handler.CampaignPost(rr, req)

	assert.NotNil(err)
}
