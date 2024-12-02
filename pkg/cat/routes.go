package cat

import (
	"vet-clinic-api/config"

	"github.com/go-chi/chi/v5"
)

func Routes(configuration *config.Config) chi.Router {
	CatConfigurator := New(configuration)
	router := chi.NewRouter()
	router.Post("/", CatConfigurator.addCatHandler)
	router.Get("/", CatConfigurator.catsHandler)
	router.Get("/{id}", CatConfigurator.catByIdHandler)
	router.Put("/edit/{id}", CatConfigurator.editCatHandler)
	router.Delete("/delete/{id}", CatConfigurator.deleteCatHandler)
	router.Get("/{id}/history", CatConfigurator.historyCatHandler)
	return router
}
