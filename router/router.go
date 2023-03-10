package router

import (
	"log"
	"net/http"
	"server/service"

	"github.com/gorilla/mux"
)

func StartRouter() {
	log.Println("Starting the HTTP server on port 8090")
	router := mux.NewRouter()

	router.HandleFunc("/users/", service.CreatePerson).Methods("POST")
	router.HandleFunc("/users/{id}", service.GetPersonByID).Methods("GET")
	router.HandleFunc("/users/", service.GetAllPerson).Methods("GET")
	router.HandleFunc("/users/{id}", service.UpdatePersonByID).Methods("PUT")
	router.HandleFunc("/users/{id}", service.DeletPersonByID).Methods("DELETE")

	router.HandleFunc("/product/", service.CreateProduct).Methods("POST")
	router.HandleFunc("/product/{id}", service.GetProductByID).Methods("GET")
	router.HandleFunc("/product/", service.GetAllProduct).Methods("GET")
	router.HandleFunc("/product/{id}", service.UpdateProductByID).Methods("PUT")
	router.HandleFunc("/product/{id}", service.DeletProductByID).Methods("DELETE")

	http.Handle("/", router)

	log.Fatal(http.ListenAndServe(":8090", router))

}
