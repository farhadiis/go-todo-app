package model

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Todo struct {
	ID        primitive.ObjectID `bson:"_id" json:"_id"`
	Title     string             `bson:"title" json:"title" binding:"required"`
	Body      string             `bson:"body" json:"body"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
}

func (t Todo) String() string {
	return fmt.Sprintf("%v (%v)", t.ID, t.Title)
}
