package router

import (
	"farhadiis/todo/interface/controller"
	"github.com/gin-gonic/gin"
)

func NewRouter(app *gin.Engine, c controller.AppController) *gin.Engine {

	app.GET("/api/v1/todo", c.GetTodos)
	app.GET("/api/v1/todo/:id", c.GetTodo)
	app.POST("/api/v1/todo", c.CreateTodo)
	app.DELETE("/api/v1/todo/:id", c.DeleteTodo)

	return app
}
