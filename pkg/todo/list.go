package todo

import "fmt"

// List all to-dos
func (t *TodoList) List() {
	for i, todo := range *t {
		i++
		fmt.Printf("%d | %s | %t \n", i, todo.Task, todo.Done)
	}
}
