package main

import (
	"server/database"
	"server/router"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// func checkIfPersonExists(personId string) bool {
// 	var person entities.Person
// 	database.Connector.First(&person, personId)
// 	if person.ID == 0 {
// 		return false
// 	}
// 	return true
// }

// func createPerson(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var person entities.Person
// 	json.NewDecoder(r.Body).Decode(&person)
// 	database.Connector.Create(&person)
// 	json.NewEncoder(w).Encode(person)
// }

// func getPersonByID(w http.ResponseWriter, r *http.Request) {
// 	personId := mux.Vars(r)["id"]
// 	if checkIfPersonExists(personId) == false {
// 		json.NewEncoder(w).Encode("Person Not Found!")
// 		return
// 	}
// 	var person entities.Person
// 	database.Connector.First(&person, personId)
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(person)
// }

// func getAllPerson(w http.ResponseWriter, r *http.Request) {
// 	var persons []entities.Person
// 	database.Connector.Find(&persons)
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(persons)
// }

// func updatePersonByID(w http.ResponseWriter, r *http.Request) {
// 	personId := mux.Vars(r)["id"]
// 	if checkIfPersonExists(personId) == false {
// 		json.NewEncoder(w).Encode("Person Not Found!")
// 		return
// 	}
// 	var person entities.Person
// 	database.Connector.First(&person, personId)
// 	json.NewDecoder(r.Body).Decode(&person)
// 	database.Connector.Save(&person)
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(person)
// }

// func deletPersonByID(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	personId := mux.Vars(r)["id"]
// 	if checkIfPersonExists(personId) == false {
// 		w.WriteHeader(http.StatusNotFound)
// 		json.NewEncoder(w).Encode("Person Not Found!")
// 		return
// 	}
// 	var person entities.Person
// 	database.Connector.Delete(&person, personId)
// 	json.NewEncoder(w).Encode("Person Deleted Successfully!")
// }

func main() {
	config :=
		database.Config{
			User:     "root",
			Password: "root",
			DB:       "gotest",
		}

	connectionString := database.GetConnectionString(config)
	database.Connect(connectionString)
	database.Migrate()

	router.StartRouter()

	// log.Println("Starting the HTTP server on port 8090")
	// router := mux.NewRouter()
	// router.HandleFunc("/create", createPerson).Methods("POST")
	// router.HandleFunc("/get/{id}", getPersonByID).Methods("GET")
	// router.HandleFunc("/get", getAllPerson).Methods("GET")
	// router.HandleFunc("/update/{id}", updatePersonByID).Methods("PUT")
	// router.HandleFunc("/delete/{id}", deletPersonByID).Methods("DELETE")
	// http.Handle("/", router)

}
