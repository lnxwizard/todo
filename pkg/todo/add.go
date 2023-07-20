package todo

import "time"

// Add a new to-do
func (t *TodoList) Add(task string) {
	todo := Item{
		Task:        task,
		Done:        false,
		Important:   false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	*t = append(*t, todo)
}

// Add a new important to-do
func (t *TodoList) AddImportant(task string) {
	todo := Item{
		Task:        task,
		Done:        false,
		Important:   true,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	*t = append(*t, todo)
}
