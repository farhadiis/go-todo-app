package presenter

import (
	"farhadiis/todo/application/presenter"
	"farhadiis/todo/domain/model"
)

type todoPresenter struct{}

func (tp *todoPresenter) ResponseTodos(todos []*model.Todo) []*model.Todo {
	for _, todo := range todos {
		todo.Title = "# " + todo.Title
	}
	return todos
}

func NewTodoPresenter() presenter.TodoPresenter {
	return &todoPresenter{}
}
