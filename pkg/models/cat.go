package models

import (
	"errors"
	"net/http"
)

type Cat struct {
	Name     string `json:"cat_name"`
	Age      int    `json:"cat_age"`
	CatBreed string `json:"cat_breed"`
	Weight   int    `json:"cat_weight"`
}

type CatHistory struct {
	Name     string `json:"cat_name"`
	Age      int    `json:"cat_age"`
	CatBreed string `json:"cat_breed"`
	Weight   int    `json:"cat_weight"`
	Visits   []Visit
}

func (c *Cat) Bind(r *http.Request) error {

	if c.Age < 0 {
		return errors.New("cat_age must be a positive integer")
	} else if c.Weight < 0 {
		return errors.New("cat_weight must be a poisitive integer")
	}

	return nil
}
