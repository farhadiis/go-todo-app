package model

import (
	"fmt"
	"time"
)

type Todo struct {
	ID    string    `json:"id"`
	Title string    `json:"title"`
	Body  string    `json:"body"`
	Date  time.Time `json:"date"`
}

func (t Todo) String() string {
	return fmt.Sprintf("%v (%v)", t.ID, t.Title)
}
