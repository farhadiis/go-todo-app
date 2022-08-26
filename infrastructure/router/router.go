package router

import (
	"farhadiis/todo/interface/controller"
	"github.com/gin-gonic/gin"
)

func NewRouter(app *gin.Engine, c controller.AppController) *gin.Engine {

	app.GET("/api/v1/todos", c.GetTodos)
	app.POST("/api/v1/todos", c.CreateTodo)

	return app
}
