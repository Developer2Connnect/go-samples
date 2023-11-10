// main.go

package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Person represents the structure of a MongoDB document
type Person struct {
	Name  string `bson:"name"`
	Age   int    `bson:"age"`
	Email string `bson:"email"`
}

func main() {
	// MongoDB connection parameters
	mongoURI := "mongodb://localhost:27017"
	dbName := "yourdbname"
	collectionName := "people"

	// Set client options
	clientOptions := options.Client().ApplyURI(mongoURI)

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	// Get a reference to the "people" collection
	collection := client.Database(dbName).Collection(collectionName)

	// Insert a document
	insertResult, err := collection.InsertOne(context.Background(), Person{"John Doe", 30, "john@example.com"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted document ID:", insertResult.InsertedID)

	// Query for documents
	var result Person
	filter := bson.M{"name": "John Doe"}
	err = collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found document: %+v\n", result)
}
