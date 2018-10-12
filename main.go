package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"phonebook-backend/models"

	"github.com/mongodb/mongo-go-driver/mongo"
)

const DBNAME = "phonebook"
const COLLECTION = "people"
const CONNECTIONSTRING = "mongodb://localhost:27017"

var people []models.Person

func init() {
	// Populates database with dummy data

	client, err := mongo.NewClient(CONNECTIONSTRING)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	db := client.Database(DBNAME)

	// Load values from JSON file to model
	byteValues, err := ioutil.ReadFile("person_data.json")
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(byteValues, &people)

	// Insert people into DB
	var ppl []interface{}
	for _, p := range people {
		ppl = append(ppl, p)
	}
	_, err = db.Collection(COLLECTION).InsertMany(context.Background(), ppl)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

}
