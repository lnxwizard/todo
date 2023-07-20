package todo

import "time"

var Todo = &TodoList{}

const TodoFile = "todos.json"

type Item struct {
	Task        string    `json:"task"`
	Done        bool      `json:"done"`
	Important   bool      `json:"important"`
	CreatedAt   time.Time `json:"createdAt"`
	CompletedAt time.Time `json:"completedAt"`
}

type TodoList []Item
