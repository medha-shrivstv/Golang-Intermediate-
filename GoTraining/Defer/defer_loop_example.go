package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("Hello")
	print()
	fmt.Println("World")
}

func print() {
	for i := 1; i <= 10; i++ {
		defer fmt.Println(i)
	}
	defer fmt.Println("Line1")
	fmt.Println("Line2")
	fmt.Println("Line3")
}
