package repository

import "farhadiis/todo/domain/model"

type TodoRepository interface {
	FindAll() ([]*model.Todo, error)
	InsertOne(*model.Todo) (string, error)
}
