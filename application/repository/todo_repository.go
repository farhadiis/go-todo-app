package repository

import "farhadiis/todo/domain/model"

type TodoRepository interface {
	FindAll() ([]*model.Todo, error)
	FindOne(string) (*model.Todo, error)
	InsertOne(*model.Todo) (string, error)
	DeleteOne(string) (int64, error)
}
