package service

import (
	"encoding/json"
	"net/http"
	"server/database"
	"server/entities"

	"github.com/gorilla/mux"
)

/*
	\brief checkIfPersonExists - function which return exist person or not
*/

func checkIfPersonExists(personId string) bool {
	var person entities.Person
	database.Connector.First(&person, personId)
	return person.ID != 0
}

/*
	\brief CreatePerson - function for router which create a new person
*/

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var person entities.Person
	json.NewDecoder(r.Body).Decode(&person)
	database.Connector.Create(&person)
	json.NewEncoder(w).Encode(person)
}

/*
	\brief GetPersonByID - function for router which return person by id
*/

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

/*
	\brief GetAllPerson - function for router which return all persons
*/

func GetAllPerson(w http.ResponseWriter, r *http.Request) {
	var persons []entities.Person
	database.Connector.Find(&persons)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(persons)
}

/*
	\brief UpdatePersonByID - function for router which update person
*/

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

/*
	\brief DeletePersonByID - function for router which delete person by id
*/

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
