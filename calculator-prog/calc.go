package main

import (
	"errors"
)

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func multi(a, b int) int {
	return a * b
}

func div(a, b int) (int, int, error) {
	var err error
	if b == 0 {
		return 0, 0, errors.New("not allowed")
	} else {
		return a / b, a % b, err
	}
}
