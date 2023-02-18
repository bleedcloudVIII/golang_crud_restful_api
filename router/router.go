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
	router.HandleFunc("/create", service.CreatePerson).Methods("POST")
	router.HandleFunc("/get/{id}", service.GetPersonByID).Methods("GET")
	router.HandleFunc("/get", service.GetAllPerson).Methods("GET")
	router.HandleFunc("/update/{id}", service.UpdatePersonByID).Methods("PUT")
	router.HandleFunc("/delete/{id}", service.DeletPersonByID).Methods("DELETE")
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8090", router))

}
