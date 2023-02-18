package service

import (
	"encoding/json"
	"net/http"
	"server/database"
	"server/entities"

	"github.com/gorilla/mux"
)

func checkIfProductExists(productId string) bool {
	var product entities.Product
	database.Connector.First(&product, productId)
	if product.ID == 0 {
		return false
	}
	return true
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product entities.Product
	json.NewDecoder(r.Body).Decode(&product)
	database.Connector.Create(&product)
	json.NewEncoder(w).Encode(product)
}

func GetProductByID(w http.ResponseWriter, r *http.Request) {
	productId := mux.Vars(r)["id"]
	if checkIfProductExists(productId) == false {
		json.NewEncoder(w).Encode("Product not found!")
		return
	}
	var product entities.Product
	database.Connector.First(&product, productId)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func GetAllProduct(w http.ResponseWriter, r *http.Request) {
	var Product []entities.Product
	database.Connector.Find(&Product)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Product)
}

func UpdateProductByID(w http.ResponseWriter, r *http.Request) {
	productId := mux.Vars(r)["id"]
	if checkIfProductExists(productId) == false {
		json.NewEncoder(w).Encode("Product not found!")
		return
	}
	var product entities.Product
	database.Connector.First(&product, productId)
	json.NewDecoder(r.Body).Decode(&product)
	database.Connector.Save(&product)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func DeletProductByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	productId := mux.Vars(r)["id"]
	if checkIfProductExists(productId) == false {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Product not found!")
		return
	}
	var product entities.Product
	database.Connector.Delete(&product, productId)
	json.NewEncoder(w).Encode("Product Deleted Successfully!")
}
