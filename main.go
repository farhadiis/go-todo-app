package main

import (
	"farhadiis/todo/infrastructure/datastore"
	"farhadiis/todo/infrastructure/router"
	"farhadiis/todo/registry"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	env := os.Getenv("GO_ENV")
	if env == "" {
		log.Fatal("You must set your 'GO_ENV' environmental variable.")
	}
	if env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
}

func main() {
	client := datastore.Connect()
	defer datastore.Disconnect(client)

	r := registry.NewRegistry(client)
	app := router.NewRouter(gin.Default(), r.NewAppController())
	app.SetTrustedProxies(nil)
	app.NoRoute(func(c *gin.Context) {
		c.String(404, "use /todos")
	})

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("You must set your 'PORT' environmental variable.")
	}
	err := app.Run("localhost:" + port)
	if err != nil {
		return
	}
}
