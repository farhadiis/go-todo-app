package controller

import (
	"farhadiis/todo/application/interactor"
	"farhadiis/todo/domain/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TodoController interface {
	GetTodos(*gin.Context)
	GetTodo(*gin.Context)
	CreateTodo(*gin.Context)
	DeleteTodo(*gin.Context)
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

func (t *todoController) GetTodo(c *gin.Context) {
	id := c.Param("id")
	todo, err := t.TodoInteractor.Get(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if todo != nil {
		c.IndentedJSON(http.StatusOK, todo)
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"_id": id, "founded": false})
	}
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

func (t *todoController) DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	ok, err := t.TodoInteractor.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if ok {
		c.IndentedJSON(http.StatusOK, gin.H{"_id": id, "deleted": ok})
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"_id": id, "deleted": ok})
	}
}

func NewTodoController(ti interactor.TodoInteractor) TodoController {
	return &todoController{ti}
}
