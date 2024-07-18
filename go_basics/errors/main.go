package main

import (
	"errors"
	"fmt"
)

func someFunc(i int) (int, error) {
	j, err := funcReturningError(i)
	if err != nil {
		return 0, fmt.Errorf("wrap error, %w", err) // error wrapping
	}
	return j, nil
}

func funcReturningError(i int) (int, error) {
	if i == 0 {
		return 0, fmt.Errorf("got %d", i)
	}
	return i, nil
}

func main() {
	i, err := someFunc(0)
	fmt.Println(i, err)
	if wrappedErr := errors.Unwrap(err); wrappedErr != nil {
		fmt.Printf("origin error: %v\n", wrappedErr)
	}

	i, err = someFunc(10)
	fmt.Println(i, err)
}
