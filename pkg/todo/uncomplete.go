package todo

import (
	"errors"
	"time"
)

func (t *TodoList) Uncomplete(index int) error {
	tdls := *t

	if index < 0 || index > len(tdls) {
		return errors.New("invalid index")
	}

	tdls[index-1].Done = false
	tdls[index-1].CompletedAt = time.Time{}

	return nil
}
