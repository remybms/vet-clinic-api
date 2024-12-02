package dbmodel

import "gorm.io/gorm"

type Visit struct {
	gorm.Model
	Date       string `json:"visit_date"`
	Reason     string `json:"visit_reason"`
	Veterinary string `json:"veterinary"`
	CatId      int    `json:"id_cat"`

	Treatments []*Treatment `gorm:"foreignKey:VisitId"`
}

type VisitRepository interface {
	Create(newVisit *Visit) (*Visit, error)
	FindById(catId string, filter string) ([]*Visit, error)
}

type visitRepository struct {
	db *gorm.DB
}

func NewVisitRepository(db *gorm.DB) VisitRepository {
	return &visitRepository{db: db}
}

func (r *visitRepository) Create(entry *Visit) (*Visit, error) {
	if err := r.db.Create(entry).Error; err != nil {
		return nil, err
	}
	return entry, nil
}

func (r *visitRepository) FindById(catId string, filter string) ([]*Visit, error) {
	var entries []*Visit
	switch filter {
	case "date":
		if err := r.db.Preload("Treatments").Where("cat_id = ?", catId).Order("date ASC").Find(&entries).Error; err != nil {
			return nil, err
		}
	case "veterinary":
		if err := r.db.Preload("Treatments").Where("cat_id = ?", catId).Order("veterinary ASC").Find(&entries).Error; err != nil {
			return nil, err
		}
	case "reason":
		if err := r.db.Preload("Treatments").Where("cat_id = ?", catId).Order("reason ASC").Find(&entries).Error; err != nil {
			return nil, err
		}
	default:
		if err := r.db.Preload("Treatments").Where("cat_id = ?", catId).Find(&entries).Error; err != nil {
			return nil, err
		}
	}
	return entries, nil
}
