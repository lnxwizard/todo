package input

import (
	"errors"
	"os"
	"strconv"
)

func GetIntegerInput() (int, error) {
	if len(os.Args) < 2 {
		return 0, errors.New("integer input not provided")
	}

	input, err := strconv.Atoi(os.Args[2])
	if err != nil {
		return 0, err
	}

	return input, nil
}
