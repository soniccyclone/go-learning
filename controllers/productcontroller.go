package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/soniccyclone/go-learning/database"
	"github.com/soniccyclone/go-learning/entities"

	"github.com/gorilla/mux"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product entities.Product
	json.NewDecoder(r.Body).Decode(&product)
	err := database.Instance.Create(&product).Error
	if err != nil {
		writeResponse(w, http.StatusInternalServerError, "Error while creating product!")
	} else {
		writeResponse(w, http.StatusCreated, product)
	}
}

func GetProductById(w http.ResponseWriter, r *http.Request) {
	productId := mux.Vars(r)["id"]
	var product entities.Product
	err := database.Instance.First(&product, productId).Error
	if err != nil {
		writeResponse(w, http.StatusInternalServerError, "Product Not Found!")
	} else {
		writeResponse(w, http.StatusOK, product)
	}
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	var products []entities.Product
	database.Instance.Find(&products)
	writeResponse(w, http.StatusOK, products)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productId := mux.Vars(r)["id"]
	var product entities.Product
	err := database.Instance.First(&product, productId).Error
	if err != nil {
		writeResponse(w, http.StatusInternalServerError, "Product Not Found!")
	} else {
		json.NewDecoder(r.Body).Decode(&product)
		database.Instance.Save(&product)
		writeResponse(w, http.StatusOK, product)
	}
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productId := mux.Vars(r)["id"]
	rows := database.Instance.Delete(&entities.Product{}, productId).RowsAffected
	if rows < 1 {
		writeResponse(w, http.StatusInternalServerError, "Product Not Found!")
	} else {
		writeResponse(w, http.StatusNoContent, "Product Deleted Successfully!")
	}
}

func writeResponse(w http.ResponseWriter, statusCode int, message interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(message)
}
