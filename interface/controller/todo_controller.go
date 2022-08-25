package controller

import (
	"farhadiis/todo/application/interactor"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TodoController interface {
	GetTodos(*gin.Context)
}

type todoController struct {
	TodoInteractor interactor.TodoInteractor
}

func (t *todoController) GetTodos(c *gin.Context) {
	todos, err := t.TodoInteractor.GetAll()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if todos == nil {
		c.IndentedJSON(http.StatusNotFound, []int{})
		return
	}
	c.IndentedJSON(http.StatusOK, todos)
}

func NewTodoController(ti interactor.TodoInteractor) TodoController {
	return &todoController{ti}
}
