package endpoints

import "github.com/igorsalves/estudos/tree/main/udemy/go-do-zero-ao-avancado/email-notifier/internal/domain/campaign"

type Handler struct {
	CampaignService campaign.Service
}
