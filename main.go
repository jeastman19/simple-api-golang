package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	ID        string   `json:"id,omitempty"`
	FirstName string   `json:"firstname,omitempty"`
	LastName  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

func GetPeople(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func GetPerson(w http.ResponseWriter, req *http.Request) {

}

func CreatePerson(w http.ResponseWriter, req *http.Request) {

}

func UpdatePerson(w http.ResponseWriter, req *http.Request) {

}

func DeletePerson(w http.ResponseWriter, req *http.Request) {

}

func main() {
	router := mux.NewRouter()

	people = append(people, Person{
		ID:        "1",
		FirstName: "Ryan",
		LastName:  "Ray",
		Address: &Address{
			City:  "Dubling",
			State: "California",
		},
	})

	people = append(people, Person{
		ID:        "2",
		FirstName: "Joe",
		LastName:  "McMillan",
	})

	// endpoints
	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/people", CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", UpdatePerson).Methods("PUT")
	router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", router))
}
