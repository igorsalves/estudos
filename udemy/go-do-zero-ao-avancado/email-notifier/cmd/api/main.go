package main

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/igorsalves/estudos/tree/main/udemy/go-do-zero-ao-avancado/email-notifier/internal/domain/campaign"
	"github.com/igorsalves/estudos/tree/main/udemy/go-do-zero-ao-avancado/email-notifier/internal/endpoints"
	"github.com/igorsalves/estudos/tree/main/udemy/go-do-zero-ao-avancado/email-notifier/internal/infrastructure/database"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(endpoints.Auth)

	db := database.NewDB()
	campaignService := campaign.ServiceImp{
		Repository: &database.CampaignRepository{Db: db},
	}
	handler := endpoints.Handler{
		CampaignService: &campaignService,
	}

	r.Get("/campaigns/{id}", endpoints.HandlerError(handler.CampaignGetById))
	r.Post("/campaigns", endpoints.HandlerError(handler.CampaignPost))
	r.Patch("/campaigns/cancel/{id}", endpoints.HandlerError(handler.CampaignCancelPatch))
	r.Delete("/campaigns/{id}", endpoints.HandlerError(handler.CampaignDelete))

	http.ListenAndServe(":3000", r)
}
