package repository

import (
	"farhadiis/todo/application/repository"
	"farhadiis/todo/domain/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type todoRepository struct {
	client *mongo.Client
}

func (t *todoRepository) FindAll() ([]*model.Todo, error) {
	return nil, nil
}

func NewTodoRepository(client *mongo.Client) repository.TodoRepository {
	return &todoRepository{client}
}
