package internal

import (
	"time"
)

type Todo struct {
	Title       string    `json:"title"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"createdAt"`
	CompletedAt time.Time `json:"completedAt"`
}

type TodoCreate struct {
	Title string
}

type TodoUpdate TodoCreate
