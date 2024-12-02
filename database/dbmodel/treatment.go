package dbmodel

import "gorm.io/gorm"

type Treatment struct {
	gorm.Model
	Medication string `json:"medication"`
	VisitId    int    `json:"id_visit"`
}

type TreatmentRepository interface {
	Create(newTreatment *Treatment) (*Treatment, error)
	FindById(idVisit string) ([]*Treatment, error)
}

type treatmentRepository struct {
	db *gorm.DB
}

func NewTreatmentRepository(db *gorm.DB) TreatmentRepository {
	return &treatmentRepository{db: db}
}

func (r *treatmentRepository) Create(entry *Treatment) (*Treatment, error) {
	if err := r.db.Create(entry).Error; err != nil {
		return nil, err
	}
	return entry, nil
}

func (r *treatmentRepository) FindById(idVisit string) ([]*Treatment, error) {
	var entries []*Treatment
	if err := r.db.Where("visit_id = ?", idVisit).Find(&entries).Error; err != nil {
		return nil, err
	}
	return entries, nil
}
