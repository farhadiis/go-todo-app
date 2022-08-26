package interactor

import (
	"farhadiis/todo/application/interactor"
	"farhadiis/todo/domain/model"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
	"time"
)

var (
	mockTodo = &model.Todo{
		ID:        primitive.NewObjectID(),
		Title:     "Go Learn",
		Body:      "Today...",
		CreatedAt: time.Time{},
	}
	mockTodos = []*model.Todo{mockTodo}
)

type (
	MockPresenter  struct{}
	MockRepository struct {
		DataStore []*model.Todo
	}
)

func (m *MockRepository) FindAll() ([]*model.Todo, error) {
	return m.DataStore, nil
}

func (m *MockRepository) FindOne(id string) (*model.Todo, error) {
	for _, todo := range m.DataStore {
		if todo.ID.Hex() == id {
			return todo, nil
		}
	}
	return nil, errors.New("todo not found")
}

func (m *MockRepository) InsertOne(todo *model.Todo) (string, error) {
	m.DataStore = append(m.DataStore, todo)
	return todo.ID.Hex(), nil
}

func (m *MockRepository) DeleteOne(id string) (int64, error) {
	for index, todo := range m.DataStore {
		if todo.ID.Hex() == id {
			m.DataStore = append(m.DataStore[:index], m.DataStore[index+1:]...)
			return 1, nil
		}
	}
	return -1, errors.New("todo not found")
}

func (m *MockPresenter) ResponseTodos(todos []*model.Todo) []*model.Todo {
	return todos
}

// TestGet calls interactor.Get, checking for return single todo.
func TestGet(t *testing.T) {
	mockRepository := &MockRepository{}
	mockPresenter := &MockPresenter{}
	in := interactor.NewTodoInteractor(mockRepository, mockPresenter)
	id, err := mockRepository.InsertOne(mockTodo)
	if id == "" || err != nil {
		t.Fatalf("Can't insert todo in repository")
	}
	todo, err := in.Get(mockTodo.ID.Hex())
	if todo != mockTodo || err != nil {
		t.Fatalf(`Get() must return %q, nill`, mockTodo)
	}
}

// TestGetAll calls interactor.GetAll, checking for return many todo.
func TestGetAll(t *testing.T) {
	mockRepository := &MockRepository{}
	mockPresenter := &MockPresenter{}
	in := interactor.NewTodoInteractor(mockRepository, mockPresenter)
	id, err := mockRepository.InsertOne(mockTodo)
	if id == "" || err != nil {
		t.Fatalf("Can't insert todo in repository")
	}
	todos, err := in.GetAll()
	if todos[0].Title != mockTodos[0].Title || todos[0].ID != mockTodos[0].ID ||
		todos[0].Body != mockTodos[0].Body || todos[0].CreatedAt != mockTodos[0].CreatedAt || err != nil {
		t.Fatalf(`GetAll return %q must equal %q, nill`, todos[0], mockTodos[0])
	}
}

// TestDelete calls interactor.Delete, checking for delete todo.
func TestDelete(t *testing.T) {
	mockRepository := &MockRepository{}
	mockPresenter := &MockPresenter{}
	in := interactor.NewTodoInteractor(mockRepository, mockPresenter)
	id, err := mockRepository.InsertOne(mockTodo)
	if id == "" || err != nil {
		t.Fatalf("can't insert todo in repository")
	}
	ok, err := in.Delete(mockTodo.ID.Hex())
	if !ok || err != nil {
		t.Fatalf(`Delete return %v must equal true, nill`, ok)
	}
}
