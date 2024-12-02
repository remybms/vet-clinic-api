package visit

import (
	"net/http"
	"vet-clinic-api/config"
	"vet-clinic-api/database/dbmodel"
	"vet-clinic-api/pkg/models"
	"vet-clinic-api/pkg/treatment"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type VisitConfigurator struct {
	*config.Config
}

func New(configuration *config.Config) *VisitConfigurator {
	return &VisitConfigurator{configuration}
}

func VisitToModel(visits []*dbmodel.Visit) []models.Visit {
	visitToModel := &models.Visit{}
	visitsEdited := []models.Visit{}
	for _, visit := range visits {
		visitToModel.Date = visit.Date
		visitToModel.Reason = visit.Reason
		visitToModel.Veterinary = visit.Veterinary
		visitToModel.CatId = visit.CatId
		visitToModel.Treatments = treatment.TreatmentToModel(visit.Treatments)
		visitsEdited = append(visitsEdited, *visitToModel)
	}
	return visitsEdited
}

func (config *VisitConfigurator) visitHistoryHandler(w http.ResponseWriter, r *http.Request) {
	catId := chi.URLParam(r, "id")
	filter := chi.URLParam(r, "filter")
	visits, err := config.VisitRepository.FindById(catId, filter)
	visitsEdited := VisitToModel(visits)
	if err != nil {
		render.JSON(w, r, map[string]string{"Error": "Failed to load all the visits"})
		return
	}
	render.JSON(w, r, visitsEdited)
}

func (config *VisitConfigurator) addVisitHandler(w http.ResponseWriter, r *http.Request) {
	req := &models.Visit{}
	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid request payload"})
		return
	}
	addVisit := &dbmodel.Visit{Date: req.Date, Reason: req.Reason, Veterinary: req.Veterinary, CatId: req.CatId}
	config.VisitRepository.Create(addVisit)
	render.JSON(w, r, map[string]string{"success": "Visit succesfully added"})

}
