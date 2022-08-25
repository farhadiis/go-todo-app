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

func NewRegistry(client *mongo.Client) Registry {
	return &registry{client}
}

func (r *registry) NewAppController() controller.AppController {
	return r.NewTodoController()
}
