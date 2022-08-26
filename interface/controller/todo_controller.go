package controller

import (
	"farhadiis/todo/application/interactor"
	"farhadiis/todo/domain/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TodoController interface {
	GetTodos(*gin.Context)
	CreateTodo(*gin.Context)
}

type todoController struct {
	TodoInteractor interactor.TodoInteractor
}

func (t *todoController) GetTodos(c *gin.Context) {
	todos, err := t.TodoInteractor.GetAll()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if todos == nil {
		c.IndentedJSON(http.StatusOK, []int{})
		return
	}
	c.IndentedJSON(http.StatusOK, todos)
}

func (t *todoController) CreateTodo(c *gin.Context) {
	var newTodo *model.Todo
	if err := c.ShouldBindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := t.TodoInteractor.Create(newTodo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, gin.H{"_id": id})
}

func NewTodoController(ti interactor.TodoInteractor) TodoController {
	return &todoController{ti}
}
