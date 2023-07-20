package todo

import (
	"errors"
	"time"
)

// Mark a to-do as completed by index
func (t *TodoList) Complete(index int) error {
	tdls := *t

	if index <= 0 || index > len(tdls) {
		return errors.New("invalid index")
	}

	tdls[index-1].CompletedAt = time.Now()
	tdls[index-1].Done = true

	return nil
}

// Mark all to-do's as completed
func (t *TodoList) CompleteAll() {
	tdls := *t

	for i := 0; i < len(*t); i++ {
		if tdls[i].Done {
			continue
		}

		tdls[i].CompletedAt = time.Now()
		tdls[i].Done = true
	}
}
