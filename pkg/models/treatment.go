package models

import (
	"errors"
	"net/http"
)

type Treatment struct {
	Medication string `json:"medication"`
	VisitId    int    `json:"id_visit"`
}

func (t *Treatment) Bind(r *http.Request) error {
	if t.VisitId < 0 {
		return errors.New("id_visit must be a positive integer")
	}
	return nil
}
