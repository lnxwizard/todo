package todo

func (t *TodoList) CountPending() int {
	totalPending := 0
	for _, todo := range *t {
		if !todo.Done {
			totalPending++
		}
	}

	return totalPending
}
