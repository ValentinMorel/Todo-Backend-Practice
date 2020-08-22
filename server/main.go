package main

import (
	"todo-app/server/models"

	"github.com/davecgh/go-spew/spew"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	id := primitive.NewObjectID()
	tasklist := models.TaskList{1, id, "test", true}
	spew.Dump(tasklist.ID)
}
