package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"../model"
)

type ProductController struct{}

func (pc ProductController) ReadAll(w http.ResponseWriter, r *http.Request) {
	m := model.ProductModel{}
	products, err := m.ReadAll()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(products)
}
