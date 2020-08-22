package middleware

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	"go.mongodb.org/mongo-driver/mongo/readpref"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	// Replace the uri with personal MongoDB deployment's connection string.
	dbname   = "test"
	collname = "todolist"
)

var collection *mongo.Collection

func init() {
	file, err := ioutil.ReadFile("cred.txt")
	if err != nil {
		log.Fatal(err)
	}
	uri := string(file)

	clientoptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientoptions)
	if err != nil {
		// If we don't connect the database, there's no point going further -> fatal.
		log.Fatal(err)
	}
	// Healthcheck.
	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB.")

	collection = client.Database(dbname).Collection(collname)
	fmt.Println(collection)
}
