package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"../model"
)

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := model.GetAllProducts()
	if err != nil {
		log.Println(err)
	}

	json.NewEncoder(w).Encode(products)
	// fmt.Fprintln(w, products)
}
