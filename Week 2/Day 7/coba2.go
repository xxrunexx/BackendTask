package main

import (
	"errors"
	"fmt"
)

func myFunc(i int) (int, error) {
	if i <= 0 {
		return -1, errors.New("Should be greater than zero")
	}
	return i, nil
}

func main() {
	result, err := myFunc(-1)
	fmt.Println(result, err)
}
