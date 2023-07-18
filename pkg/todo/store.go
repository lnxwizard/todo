package todo

import (
	"encoding/json"
	"os"
)

// Store all to-dos in json file
func (t *TodoList) Store(filename string) error {
	data, err := json.MarshalIndent(t, "", "    ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}
