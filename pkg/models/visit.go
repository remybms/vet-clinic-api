package models

import (
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

	return nil
}
