package registry

import (
	"farhadiis/todo/interface/controller"
	"go.mongodb.org/mongo-driver/mongo"
)

type Registry interface {
	NewAppController() controller.AppController
}

type registry struct {
	MongoClient *mongo.Client
}

func NewRegistry(mongoClient *mongo.Client) Registry {
	return &registry{mongoClient}
}

func (r *registry) NewAppController() controller.AppController {
	return r.NewTodoController()
}
