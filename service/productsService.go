package service

import (
	"encoding/json"
	"net/http"
	"server/database"
	"server/entities"

	"github.com/gorilla/mux"
)

/*
	\brief checkIfProductsExists - check exists product or not
*/

func checkIfProductExists(productId string) bool {
	var product entities.Product
	database.Connector.First(&product, productId)
	return product.ID != 0
}

/*
	\brief CreateProducts - function for router which create a new product
*/

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var NewProduct entities.Product
	json.NewDecoder(r.Body).Decode(&NewProduct)
	database.Connector.Create(&NewProduct)
	json.NewEncoder(w).Encode(NewProduct)
}

/*
	\brief GetProductById - function for router which return product by id
*/

func GetProductByID(w http.ResponseWriter, r *http.Request) {
	productId := mux.Vars(r)["id"]
	if !checkIfProductExists(productId) {
		json.NewEncoder(w).Encode("Product not found!")
		return
	}
	var product entities.Product
	database.Connector.First(&product, productId)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

/*
	\brief GetAllProduct - function for router which return all products
*/

func GetAllProduct(w http.ResponseWriter, r *http.Request) {
	var product []entities.Product
	database.Connector.Find(&product)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

/*
	\brief CreateProducts - function for router which update a product
*/

func UpdateProductByID(w http.ResponseWriter, r *http.Request) {
	productId := mux.Vars(r)["id"]
	if !checkIfProductExists(productId) {
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

/*
	\brief DeleteProductByID - function for router which delete product by id
*/

func DeletProductByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	productId := mux.Vars(r)["id"]
	if !checkIfProductExists(productId) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Product not found!")
		return
	}
	var product entities.Product
	database.Connector.Delete(&product, productId)
	json.NewEncoder(w).Encode("Product Deleted Successfully!")
}
