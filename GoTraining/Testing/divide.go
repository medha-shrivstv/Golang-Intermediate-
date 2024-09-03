package main

import (
	"errors"
	"fmt"
	"strconv"
)

// Divide function according to Go's approach
func Dividing(n1, n2 string) (int, error) {
	no1, err := strconv.Atoi(n1)
	if err != nil {
		return 0, fmt.Errorf("InvalidArg1: %v", err)
	}

	no2, err := strconv.Atoi(n2)
	if err != nil {
		return 0, fmt.Errorf("InvalidArg2: %v", err)
	}

	if no2 == 0 {
		return 0, errors.New("Arg2 is zero")
	}

	return no1 / no2, nil
}

func main() {
	result, err := Dividing("10", "2")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}
