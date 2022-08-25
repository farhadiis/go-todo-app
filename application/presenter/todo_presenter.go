package presenter

import "farhadiis/todo/domain/model"

type TodoPresenter interface {
	ResponseTodos([]*model.Todo) []*model.Todo
}
