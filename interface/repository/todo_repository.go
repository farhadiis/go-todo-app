package repository

import (
	"context"
	"farhadiis/todo/application/repository"
	"farhadiis/todo/domain/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type todoRepository struct {
	client *mongo.Client
}

func (t *todoRepository) FindAll() ([]*model.Todo, error) {
	coll := t.client.Database("todo").Collection("todos")
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
	return &todoRepository{client}
}
