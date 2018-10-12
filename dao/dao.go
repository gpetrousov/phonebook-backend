package dao

import (
	"context"
	"log"
	"phonebook-backend/models"

	"github.com/mongodb/mongo-go-driver/mongo"
)

// PersonDAO allows us to create DAO element
type PersonDAO struct {
	Server     string
	Database   string
	Collection string
}

var db *mongo.Database

// Connect establish a connection to database
func (c *PersonDAO) Connect() {
	client, err := mongo.NewClient("mongodb://mongo:27017")
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	// Collection types can be used to access the database
	db = client.Database("baz")
}

// InsertManyValues inserts many items from byte slice
func (c *PersonDAO) InsertManyValues(passengers []models.Person) {
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
func (c *PersonDAO) GetAllPassengers() []models.Person {
	cur, err := db.Collection("titanic").Find(context.Background(), nil, nil)
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
func (c *PersonDAO) DeletePassenger(passenger models.Person) error {
	_, err := db.Collection("titanic").DeleteMany(context.Background(), passenger, nil)
	return err
}
