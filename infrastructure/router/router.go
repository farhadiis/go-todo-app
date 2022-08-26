package router

import (
	"farhadiis/todo/interface/controller"
	"github.com/gin-gonic/gin"
)

func NewRouter(e *gin.Engine, c controller.AppController) *gin.Engine {

	e.GET("/api/v1/todos", c.GetTodos)
	e.POST("/api/v1/todos", c.CreateTodo)

	return e
}
