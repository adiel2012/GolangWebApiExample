package controllers

import (
	"encoding/json"
	"gowebapi/model"
	"net/http"
)

var productList = []model.Product{model.Product{Id: "1", Name: "Adiel", Age: 37}}

func ListProducts(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(productList)
}

func AddProduct(w http.ResponseWriter, r *http.Request) {
	//productList = append(productList, model.Product{Id: "2", Name: "Emi", Age: 32})
}
