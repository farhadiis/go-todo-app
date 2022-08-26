package repository

import (
	"context"
	"farhadiis/todo/application/repository"
	"farhadiis/todo/domain/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type todoRepository struct {
	db *mongo.Database
}

func (t *todoRepository) FindAll() ([]*model.Todo, error) {
	coll := t.db.Collection("todos")
	cursor, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var todos []*model.Todo
	err = cursor.All(context.TODO(), &todos)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return todos, nil
}

func (t *todoRepository) FindOne(id string) (*model.Todo, error) {
	coll := t.db.Collection("todos")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Invalid id")
	}
	filter := bson.D{{"_id", objectId}}
	var todo *model.Todo
	err = coll.FindOne(context.TODO(), filter).Decode(&todo)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return todo, nil
}

func (t *todoRepository) InsertOne(todo *model.Todo) (string, error) {
	coll := t.db.Collection("todos")
	doc := bson.D{{"title", todo.Title}, {"body", todo.Body}}
	result, err := coll.InsertOne(context.TODO(), doc)
	if err != nil {
		log.Println(err)
		return "", err
	}
	id := result.InsertedID.(primitive.ObjectID)
	return id.Hex(), nil
}

func (t *todoRepository) DeleteOne(id string) (int64, error) {
	coll := t.db.Collection("todos")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Invalid id")
	}
	filter := bson.D{{"_id", objectId}}
	result, err := coll.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Println(err)
		return -1, err
	}
	return result.DeletedCount, nil
}

func NewTodoRepository(client *mongo.Client) repository.TodoRepository {
	return &todoRepository{client.Database("todo")}
}
