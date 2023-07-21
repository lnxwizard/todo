package todo

import (
	"errors"
)

func (t *TodoList) MarkImportant(index int) error {
	tdls := *t

	if index <= 0 || index > len(tdls) {
		return errors.New("invalid index")
	}

	tdls[index-1].Important = true

	return nil
}
