package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Ensure two arguments are passed
	if len(os.Args) < 3 {
		fmt.Println("Error: Please provide two integers as arguments.")
		return
	}

	// Access the command-line arguments
	arg1 := os.Args[1]
	arg2 := os.Args[2]

	// Convert the arguments from string to int
	num1, err1 := strconv.Atoi(arg1)
	if err1 != nil {
		fmt.Println("Error: The first argument is not a valid integer.")
		return
	}

	num2, err2 := strconv.Atoi(arg2)
	if err2 != nil {
		fmt.Println("Error: The second argument is not a valid integer.")
		return
	}

	// Calculate the sum
	sum := num1 + num2

	// Display the sum
	fmt.Println("Sum:", sum)
}
