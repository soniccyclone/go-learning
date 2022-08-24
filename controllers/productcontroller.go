package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/soniccyclone/go-learning/entities"
	"gorm.io/gorm"

	"github.com/gorilla/mux"
)

type ProductController struct {
	Database *gorm.DB
}

func (c *ProductController) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product entities.Product
	json.NewDecoder(r.Body).Decode(&product)
	err := c.Database.Create(&product).Error
	if err != nil {
		writeResponse(w, http.StatusInternalServerError, "Error while creating product!")
	} else {
		writeResponse(w, http.StatusCreated, product)
	}
}

func (c *ProductController) GetProductById(w http.ResponseWriter, r *http.Request) {
	productId := mux.Vars(r)["id"]
	var product entities.Product
	err := c.Database.First(&product, productId).Error
	if err != nil {
		writeResponse(w, http.StatusInternalServerError, "Product Not Found!")
	} else {
		writeResponse(w, http.StatusOK, product)
	}
}

func (c *ProductController) GetProducts(w http.ResponseWriter, r *http.Request) {
	var products []entities.Product
	c.Database.Find(&products)
	writeResponse(w, http.StatusOK, products)
}

func (c *ProductController) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productId := mux.Vars(r)["id"]
	var product entities.Product
	err := c.Database.First(&product, productId).Error
	if err != nil {
		writeResponse(w, http.StatusInternalServerError, "Product Not Found!")
	} else {
		json.NewDecoder(r.Body).Decode(&product)
		c.Database.Save(&product)
		writeResponse(w, http.StatusOK, product)
	}
}

func (c *ProductController) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productId := mux.Vars(r)["id"]
	rows := c.Database.Delete(&entities.Product{}, productId).RowsAffected
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
