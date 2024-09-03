package main

import (
	"fmt"
)

func main() {
	num1 := 10
	num2 := 5
	sum, subtraction := calculate(num1, num2)
	fmt.Println("Sum:", sum)
	fmt.Println("Subtraction:", subtraction)
}
