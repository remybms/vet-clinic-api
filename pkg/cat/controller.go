package cat

import (
	"net/http"
	"vet-clinic-api/config"
	"vet-clinic-api/database/dbmodel"
	"vet-clinic-api/pkg/models"
	"vet-clinic-api/pkg/visit"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type CatConfigurator struct {
	*config.Config
}

func New(configuration *config.Config) *CatConfigurator {
	return &CatConfigurator{configuration}
}

func catsToModel(cats []*dbmodel.Cat) []models.Cat {
	catsToModel := &models.Cat{}
	catsEdited := []models.Cat{}
	for _, cat := range cats {
		catsToModel.Name = cat.Name
		catsToModel.Age = cat.Age
		catsToModel.CatBreed = cat.CatBreed
		catsToModel.Weight = cat.Weight
		catsEdited = append(catsEdited, *catsToModel)
	}
	return catsEdited
}

func (config *CatConfigurator) catsHandler(w http.ResponseWriter, r *http.Request) {
	cats, err := config.CatRepository.FindAll()
	catsEdited := catsToModel(cats)
	if err != nil {
		render.JSON(w, r, map[string]string{"Error": "Failed to load all the cats"})
		return
	}
	render.JSON(w, r, catsEdited)
}

func (config *CatConfigurator) catByIdHandler(w http.ResponseWriter, r *http.Request) {
	catId := chi.URLParam(r, "id")
	cats, err := config.CatRepository.FindById(catId)
	catEdited := catsToModel(cats)
	if err != nil {
		render.JSON(w, r, map[string]string{"Error": "Failed to load the wanted cat"})
		return
	}
	render.JSON(w, r, catEdited)
}

func (config *CatConfigurator) addCatHandler(w http.ResponseWriter, r *http.Request) {
	req := &models.Cat{}
	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}
	addCat := &dbmodel.Cat{Name: req.Name, Age: req.Age, CatBreed: req.CatBreed, Weight: req.Weight}
	config.CatRepository.Create(addCat)
	render.JSON(w, r, map[string]string{"success": "New cat successfully added"})
}

func (config *CatConfigurator) editCatHandler(w http.ResponseWriter, r *http.Request) {
	req := &models.Cat{}
	catId := chi.URLParam(r, "id")
	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}
	updatedCat := &dbmodel.Cat{Name: req.Name, Age: req.Age, CatBreed: req.CatBreed, Weight: req.Weight}
	config.CatRepository.Update(updatedCat, catId)
	render.JSON(w, r, map[string]string{"success": "Cat successfully updated"})
}

func (config *CatConfigurator) deleteCatHandler(w http.ResponseWriter, r *http.Request) {
	catId := chi.URLParam(r, "id")
	cat, err := config.CatRepository.FindById(catId)
	if err != nil {
		render.JSON(w, r, map[string]string{"Error": "Failed to find the wanted cat"})
		return
	}
	config.CatRepository.Delete(cat[0])
	render.JSON(w, r, map[string]string{"success": "Cat successfully deleted"})
}

func (config *CatConfigurator) historyCatHandler(w http.ResponseWriter, r *http.Request) {
	catId := chi.URLParam(r, "id")
	cats, err := config.CatRepository.History(catId)
	catToModel := &models.CatHistory{}
	catsEdited := []models.CatHistory{}
	for _, cat := range cats {
		catToModel.Name = cat.Name
		catToModel.Age = cat.Age
		catToModel.CatBreed = cat.CatBreed
		catToModel.Weight = cat.Weight
		catToModel.Visits = visit.VisitToModel(cat.Visits)
		catsEdited = append(catsEdited, *catToModel)
	}
	if err != nil {
		render.JSON(w, r, map[string]string{"Error": "Failed to find the wanted cat"})
		return
	}
	render.JSON(w, r, catsEdited)
}
