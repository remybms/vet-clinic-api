package dbmodel

import "gorm.io/gorm"

type Cat struct {
	gorm.Model
	Name     string `json:"cat_name"`
	Age      int    `json:"cat_age"`
	CatBreed string `json:"cat_breed"`
	Weight   int    `json:"cat_weight"`

	Visits     []*Visit     `gorm:"foreignKey:CatId"`
}

type CatRepository interface {
	Create(newCat *Cat) (*Cat, error)
	FindAll() ([]*Cat, error)
	FindById(catId string) ([]*Cat, error)
	Delete(catToDelete *Cat) error
	Update(catToUpdate *Cat, catId string) error
	History(catId string) ([]*Cat, error)
}

type catRepository struct {
	db *gorm.DB
}

func NewCatRepository(db *gorm.DB) CatRepository {
	return &catRepository{db: db}
}

func (r *catRepository) Create(cat *Cat) (*Cat, error) {
	if err := r.db.Create(cat).Error; err != nil {
		return nil, err
	}
	return cat, nil
}

func (r *catRepository) Delete(catToDelete *Cat) error {
	if err := r.db.Delete(catToDelete).Error; err != nil {
		return err
	}
	return nil
}

func (r *catRepository) FindAll() ([]*Cat, error) {
	var cats []*Cat
	if err := r.db.Find(&cats).Error; err != nil {
		return nil, err
	}
	return cats, nil
}

func (r *catRepository) FindById(catId string) ([]*Cat, error) {
	var entry []*Cat
	if err := r.db.Where("id = ?", catId).Find(&entry).Error; err != nil {
		return nil, err
	}
	return entry, nil
}

func (r *catRepository) Update(catToUpdate *Cat, catId string) error {
	if err := r.db.Where("id = ?", catId).Updates(catToUpdate).Error; err != nil {
		return err
	}
	return nil
}

func (r *catRepository) History(catId string) ([]*Cat, error) {
	var entry []*Cat
	if err := r.db.Preload("Visits.Treatments").Where("id = ?", catId).Find(&entry).Error; err != nil {
		return nil, err
	}
	return entry, nil
}
