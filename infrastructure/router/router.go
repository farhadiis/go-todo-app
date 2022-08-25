package router

import (
	"farhadiis/todo/interface/controller"
	"github.com/gin-gonic/gin"
)

func NewRouter(e *gin.Engine, c controller.AppController) *gin.Engine {

	e.GET("/todos", c.GetTodos)

	return e
}
