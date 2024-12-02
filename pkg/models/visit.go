package models

import (
	"errors"
	"net/http"
)

type Visit struct {
	Date       string `json:"visit_date"`
	Reason     string `json:"visit_reason"`
	Veterinary string `json:"veterinary"`
	CatId      int    `json:"id_cat"`

	Treatments []Treatment
}

func (v *Visit) Bind(r *http.Request) error {
	if v.CatId < 0 {
		return errors.New("id_cat must be a positive integer")
	}
	return nil
}
