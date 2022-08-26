package main

import (
	"farhadiis/todo/infrastructure/datastore"
	"farhadiis/todo/infrastructure/router"
	"farhadiis/todo/registry"
	"farhadiis/todo/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	mode := utils.GetEnv("MODE")
	if mode == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
}

func main() {
	mongoClient := datastore.Connect()
	defer datastore.Disconnect(mongoClient)

	reg := registry.NewRegistry(mongoClient)
	app := router.NewRouter(gin.Default(), reg.NewAppController())
	app.SetTrustedProxies(nil)

	port := utils.GetEnv("PORT")
	err := app.Run("localhost:" + port)
	if err != nil {
		return
	}
}
