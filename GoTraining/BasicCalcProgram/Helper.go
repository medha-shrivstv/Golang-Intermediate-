package main

import "fmt"

func calculate(a, b int) (int, int) {
	fmt.Println("Helper.go - calculate function")
	sum := a + b
	sub := a - b
	return sum, sub
}
