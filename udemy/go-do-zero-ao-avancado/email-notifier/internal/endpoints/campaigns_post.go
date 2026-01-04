package endpoints

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/igorsalves/estudos/tree/main/udemy/go-do-zero-ao-avancado/email-notifier/internal/domain/campaign"
)

func (h *Handler) CampaignPost(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	var request campaign.NewCampaignRequest
	render.DecodeJSON(r.Body, &request)
	email := r.Context().Value("email").(string)
	request.CreatedBy = email

	id, err := h.CampaignService.Create(request)

	return map[string]string{"id": id}, 201, err
}
