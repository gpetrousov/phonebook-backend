package handlers

import (
	"encoding/json"
	"net/http"
	"phonebook-backend/models"

	"github.com/gorilla/mux"
)

var people []models.Person

// GetPersonEndpoint gets a person
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

// GetPeopleEndpoint gets all lpeople
func GetPeopleEndpoint(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}

// CreatePersonEndpoint creta a person
func CreatePersonEndpoint(w http.ResponseWriter, r *http.Request) {
	var person models.Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	people = append(people, person)
	json.NewEncoder(w).Encode(person)
}

// DeletePersonEndpoint delets a person
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

// UpdatePersonEndpoint updates a person
func UpdatePersonEndpoint(w http.ResponseWriter, r *http.Request) {
	var person models.Person
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
