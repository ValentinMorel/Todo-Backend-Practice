package main

import (
	"todo-app/server/models"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	id := primitive.NewObjectID()
	tasklist := models.TaskList{1, id, "test", true}
	spew.Dump(tasklist)
}
