package models

import "net/http"

type Treatment struct {
	Medication string `json:"medication"`
	VisitId int `json:"id_visit"`
}

func (t *Treatment) Bind(r *http.Request) error {

	return nil
}
