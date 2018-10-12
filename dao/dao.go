package dao

import (
	"context"
	"log"
	"phonebook-backend/models"

	"github.com/mongodb/mongo-go-driver/mongo"
)

// CONNECTIONSTRING DB connection string
const CONNECTIONSTRING = "mongodb://localhost:27017"

// DBNAME Database name
const DBNAME = "phonebook"

var db *mongo.Database

// Connect establish a connection to database
func init() {
	client, err := mongo.NewClient(CONNECTIONSTRING)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	// Collection types can be used to access the database
	db = client.Database(DBNAME)
}

// InsertManyValues inserts many items from byte slice
func InsertManyValues(passengers []models.Person) {
	var ppl []interface{}
	for _, p := range passengers {
		ppl = append(ppl, p)
	}
	_, err := db.Collection("titanic").InsertMany(context.Background(), ppl)
	if err != nil {
		log.Fatal(err)
	}
}

// GetAllPassengers returns all passengers from DB
func GetAllPassengers() []models.Person {
	cur, err := db.Collection("people").Find(context.Background(), nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	var elements []models.Person
	var elem models.Person
	// Get the next result from the cursor
	for cur.Next(context.Background()) {
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		elements = append(elements, elem)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	cur.Close(context.Background())
	return elements
}

// DeletePassenger deletes an existing passenger
func DeletePassenger(passenger models.Person) error {
	_, err := db.Collection("titanic").DeleteMany(context.Background(), passenger, nil)
	return err
}
