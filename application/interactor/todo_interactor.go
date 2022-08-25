package interactor

import (
	"farhadiis/todo/application/presenter"
	"farhadiis/todo/application/repository"
	"farhadiis/todo/domain/model"
)

type TodoInteractor interface {
	GetAll() ([]*model.Todo, error)
}

type todoInteractor struct {
	TodoRepository repository.TodoRepository
	TodoPresenter  presenter.TodoPresenter
}

func (t *todoInteractor) GetAll() ([]*model.Todo, error) {
	todos, err := t.TodoRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return t.TodoPresenter.ResponseTodos(todos), nil
}

func NewTodoInteractor(r repository.TodoRepository, p presenter.TodoPresenter) TodoInteractor {
	return &todoInteractor{r, p}
}
