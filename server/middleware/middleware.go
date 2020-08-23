package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"todo-app/server/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

// Get task route
func GetAllTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	payload := getalltask()
	json.NewEncoder(w).Encode(payload)
}

// Get all tasks from DB
func getalltask() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var results []primitive.M
	for cur.Next(context.Background()) {
		var result bson.M
		e := cur.Decode(&result)
		if e != nil {
			log.Fatal(e)
		}
		results = append(results, result)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	cur.Close(context.Background())
	return results
}

func insertonetask(task models.TaskList) {
	insertresult, err := collection.InsertOne(context.Background(), task)

	if err != nil {
		log.Fatal(err)
	}
	log.Println("Inserted a record : ", insertresult.InsertedID)
}

func deleteonetask(task string) {
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}
	d, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Deleted Document : ", d.DeletedCount)
}

func main() {
	//id := primitive.NewObjectID()

	//tasklist := models.TaskList{1, id, "test", true}
	//insertonetask(tasklist)
	//results := getalltask()
	//fmt.Println(results)
	deleteonetask("5f42934eaf2fb05a5599702a")

}
