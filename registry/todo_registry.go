package registry

import (
	"farhadiis/todo/application/interactor"
	p "farhadiis/todo/application/presenter"
	r "farhadiis/todo/application/repository"
	"farhadiis/todo/interface/controller"
	"farhadiis/todo/interface/presenter"
	"farhadiis/todo/interface/repository"
)

func (r *registry) NewTodoController() controller.TodoController {
	return controller.NewTodoController(r.NewTodoInteractor())
}

func (r *registry) NewTodoInteractor() interactor.TodoInteractor {
	return interactor.NewTodoInteractor(r.NewTodoRepository(), r.NewTodoPresenter())
}

func (r *registry) NewTodoRepository() r.TodoRepository {
	return repository.NewTodoRepository(r.MongoClient)
}

func (r *registry) NewTodoPresenter() p.TodoPresenter {
	return presenter.NewTodoPresenter()
}
