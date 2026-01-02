package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/igorsalves/estudos/tree/main/udemy/go-do-zero-ao-avancado/email-notifier/internal/domain/campaign"
	"github.com/igorsalves/estudos/tree/main/udemy/go-do-zero-ao-avancado/email-notifier/internal/endpoints"
	"github.com/igorsalves/estudos/tree/main/udemy/go-do-zero-ao-avancado/email-notifier/internal/infrastructure/database"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	db := database.NewDB()
	campaignService := campaign.ServiceImp{
		Repository: &database.CampaignRepository{Db: db},
	}
	handler := endpoints.Handler{
		CampaignService: &campaignService,
	}

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	r.Route("/campaigns", func(r chi.Router) {
		r.Use(endpoints.Auth)

		r.Get("/{id}", endpoints.HandlerError(handler.CampaignGetById))
		r.Post("/", endpoints.HandlerError(handler.CampaignPost))
		r.Patch("/cancel/{id}", endpoints.HandlerError(handler.CampaignCancelPatch))
		r.Delete("/{id}", endpoints.HandlerError(handler.CampaignDelete))
	})

	http.ListenAndServe(":3000", r)
}
