package treatment

import (
	"net/http"
	"vet-clinic-api/config"
	"vet-clinic-api/database/dbmodel"
	"vet-clinic-api/pkg/models"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type TreatmentConfigurator struct {
	*config.Config
}

func New(configuration *config.Config) *TreatmentConfigurator {
	return &TreatmentConfigurator{configuration}
}

func TreatmentToModel(treatments []*dbmodel.Treatment) []models.Treatment {
	treatmentToModel := &models.Treatment{}
	treatmentsEdited := []models.Treatment{}
	for _, treatment := range treatments {
		treatmentToModel.Medication = treatment.Medication
		treatmentToModel.VisitId = treatment.VisitId
		treatmentsEdited = append(treatmentsEdited, *treatmentToModel)
	}
	return treatmentsEdited
}

func (config *TreatmentConfigurator) treatmentHandler(w http.ResponseWriter, r *http.Request) {
	visitId := chi.URLParam(r, "id")
	treatments, err := config.TreatmentRepository.FindById(visitId)
	treatmentsEdited := TreatmentToModel(treatments)
	if err != nil {
		render.JSON(w, r, map[string]string{"Error": "Failed to load all the visits"})
		return
	}
	render.JSON(w, r, treatmentsEdited)
}

func (config *TreatmentConfigurator) addTreatmentHandler(w http.ResponseWriter, r *http.Request) {
	req := &models.Treatment{}
	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid request payload"})
		return
	}
	addTreatment := &dbmodel.Treatment{Medication: req.Medication, VisitId: req.VisitId}
	config.TreatmentRepository.Create(addTreatment)
	render.JSON(w, r, map[string]string{"success": "Treatment succesfully added"})
}
