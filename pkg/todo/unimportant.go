package todo

import "errors"

func (t *TodoList) Unimportant(index int) error {
	tdls := *t

	if index < 0 || index > len(tdls) {
		return errors.New("invalid index")
	}

	tdls[index-1].Important = false
	
	return nil
}
