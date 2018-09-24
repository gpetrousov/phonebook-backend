package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	ID          string `json:"id,omitempty"`
	Firstname   string `json:"firstname,omitempty"`
	Lastname    string `json:"lastname,omitempty"`
	Contactinfo `json:"contactinfo,omitempty"`
}
type Contactinfo struct {
	City    string `json:"city,omitempty"`
	Zipcode string `json:"Zipcode,omitempty"`
	Phone   string `json:"phone,omitempty"`
}

var people []Person

func GetPersonEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, p := range people {
		if p.ID == params["id"] {
			json.NewEncoder(w).Encode(p)
			return
		}
	}
	json.NewEncoder(w).Encode("Person not found")
}

func GetPeopleEndpoint(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func CreatePersonEndpoint(w http.ResponseWriter, r *http.Request) {
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	people = append(people, person)
	json.NewEncoder(w).Encode(person)
}

func DeletePersonEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for i, p := range people {
		if p.ID == params["id"] {
			copy(people[i:], people[i+1:])
			people = people[:len(people)-1]
			break
		}
	}
	json.NewEncoder(w).Encode(people)
}

func UpdatePersonEndpoint(w http.ResponseWriter, r *http.Request) {
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	params := mux.Vars(r)
	for i, p := range people {
		if p.ID == params["id"] {
			people[i] = person
			json.NewEncoder(w).Encode(person)
			break
		}
	}
}

func main() {
	router := mux.NewRouter()
	people = append(people, Person{ID: "1", Firstname: "Bruce", Lastname: "Wayne", Contactinfo: Contactinfo{City: "Gotham", Zipcode: "735", Phone: "012345678"}})
	people = append(people, Person{ID: "2", Firstname: "Clark", Lastname: "Kent", Contactinfo: Contactinfo{City: "Metropolis", Zipcode: "62960", Phone: "9876543210"}})
	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/people", CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePersonEndpoint).Methods("DELETE")
	router.HandleFunc("/people/{id}", UpdatePersonEndpoint).Methods("PUT")
	fmt.Println("Starting server on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", router))
}
