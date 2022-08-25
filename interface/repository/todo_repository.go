package repository

import (
	"context"
	"farhadiis/todo/application/repository"
	"farhadiis/todo/domain/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type todoRepository struct {
	db *mongo.Database
}

func (t *todoRepository) FindAll() ([]*model.Todo, error) {
	coll := t.db.Collection("todos")
	cursor, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}
	var todos []*model.Todo
	err = cursor.All(context.TODO(), &todos)
	if err != nil {
		panic(err)
	}
	return todos, err
}

func NewTodoRepository(client *mongo.Client) repository.TodoRepository {
	return &todoRepository{client.Database("todo")}
}
