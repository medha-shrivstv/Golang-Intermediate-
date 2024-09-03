package main

import (
	"fmt"
)

func main() {
	no1, no2 := 0, 0
	fmt.Println("Enter two numbers:")
	fmt.Scan(&no1)
	fmt.Scan(&no2)

	fmt.Println("Add:", add(no1, no2))
	fmt.Println("Sub:", sub(no1, no2))
	fmt.Println("Mult:", mult(no1, no2))
	fmt.Println("Divide:", divide(no1, no2))
}

func add(a, b int) (result int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in add:", r)
		}
	}()
	return a + b
}

func sub(a, b int) (result int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in sub:", r)
		}
	}()
	return a - b
}

func mult(a, b int) (result int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in mult:", r)
		}
	}()
	return a * b
}

func divide(a, b int) (result int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in divide:", r)
		}
	}()
	if b == 0 {
		panic("division by zero")
	}
	return a / b
}
