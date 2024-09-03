package main

import (
	"errors"
	"fmt"
	"strconv"
)

func Add(n1, n2 int) int {
	return n1 + n2
}

func Divide(n1, n2 string) (int, error) {
	no1, _ := strconv.Atoi(n1)
	no2, _ := strconv.Atoi(n2)

	if no2 == 0 {
		return 0, errors.New("division by zero")
	}
	return no1 / no2, nil
}

func main() {
	fmt.Println("Add = ", Add(10, 10))
	result, err := Divide("100", "10")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Divide = ", result)
	}
}
