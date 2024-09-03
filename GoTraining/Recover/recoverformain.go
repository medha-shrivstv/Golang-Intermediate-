package main

import (
	"fmt"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main:", r)
		}
	}()

	no1, no2 := 0, 0
	fmt.Println("Enter two numbers:")
	fmt.Scan(&no1)
	fmt.Scan(&no2)

	fmt.Println("Add:", add(no1, no2))
	fmt.Println("Sub:", sub(no1, no2))
	fmt.Println("Mult:", mult(no1, no2))
	fmt.Println("Divide:", divide(no1, no2))
}

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func mult(a, b int) int {
	return a * b
}

func divide(a, b int) int {
	if b == 0 {
		panic("division by zero")
	}
	return a / b
}
