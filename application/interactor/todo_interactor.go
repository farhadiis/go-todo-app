package interactor

import (
	"farhadiis/todo/application/presenter"
	"farhadiis/todo/application/repository"
	"farhadiis/todo/domain/model"
)

type TodoInteractor interface {
	GetAll() ([]*model.Todo, error)
	Create(*model.Todo) (string, error)
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

func (t *todoInteractor) Create(todo *model.Todo) (string, error) {
	id, err := t.TodoRepository.InsertOne(todo)
	if err != nil {
		return "", err
	}
	return id, nil
}

func NewTodoInteractor(r repository.TodoRepository, p presenter.TodoPresenter) TodoInteractor {
	return &todoInteractor{r, p}
}
