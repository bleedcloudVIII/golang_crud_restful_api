package service

import (
	"encoding/json"
	"net/http"
	"server/database"
	"server/entities"

	"github.com/gorilla/mux"
)

func checkIfPersonExists(personId string) bool {
	var person entities.Person
	database.Connector.First(&person, personId)
	if person.ID == 0 {
		return false
	}
	return true
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var person entities.Person
	json.NewDecoder(r.Body).Decode(&person)
	database.Connector.Create(&person)
	json.NewEncoder(w).Encode(person)
}

func GetPersonByID(w http.ResponseWriter, r *http.Request) {
	personId := mux.Vars(r)["id"]
	if !checkIfPersonExists(personId) {
		json.NewEncoder(w).Encode("Person Not Found!")
		return
	}
	var person entities.Person
	database.Connector.First(&person, personId)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(person)
}

func GetAllPerson(w http.ResponseWriter, r *http.Request) {
	var persons []entities.Person
	database.Connector.Find(&persons)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(persons)
}

func UpdatePersonByID(w http.ResponseWriter, r *http.Request) {
	personId := mux.Vars(r)["id"]
	if !checkIfPersonExists(personId) {
		json.NewEncoder(w).Encode("Person Not Found!")
		return
	}
	var person entities.Person
	database.Connector.First(&person, personId)
	json.NewDecoder(r.Body).Decode(&person)
	database.Connector.Save(&person)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(person)
}

func DeletPersonByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	personId := mux.Vars(r)["id"]
	if !checkIfPersonExists(personId) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Person Not Found!")
		return
	}
	var person entities.Person
	database.Connector.Delete(&person, personId)
	json.NewEncoder(w).Encode("Person Deleted Successfully!")
}
