package visit

import (
	"vet-clinic-api/config"

	"github.com/go-chi/chi/v5"
)

func Routes(configuration *config.Config) chi.Router {
	VisitConfigurator := New(configuration)
	router := chi.NewRouter()
	router.Post("/", VisitConfigurator.addVisitHandler)
	router.Get("/{id}", VisitConfigurator.visitHistoryHandler)
	router.Get("/{id}/{filter}", VisitConfigurator.visitHistoryHandler)
	return router
}
