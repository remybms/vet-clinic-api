package treatment

import (
	"vet-clinic-api/config"

	"github.com/go-chi/chi/v5"
)

func Routes(configuration *config.Config) chi.Router {
	TreatmentConfigurator := New(configuration)
	router := chi.NewRouter()
	router.Post("/", TreatmentConfigurator.addTreatmentHandler)
	router.Get("/{id}", TreatmentConfigurator.treatmentHandler)

	return router
}
