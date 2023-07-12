package todo

import "errors"

// Delete a to-do by index
func (t *TodoList) Delete(index int) error {
	tdls := *t

	if index <= 0 || index > len(tdls) {
		return errors.New("invalid index")
	}

	*t = append(tdls[:index-1], tdls[index:]...)

	return nil
}
