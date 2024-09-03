package main

import "fmt"

func main() {
	no1 := process(add)
	fmt.Println(no1)
}
func add(a, b int) int {
	fmt.Println("Lab4.go - add function ")
	return a + b
}

func process(fn func(a, b int) int) int {
	fmt.Println("process in invoked with function as parameter ")
	// this process function is common for accepting two inputs and giving output
	return fn(1000, 100)
}
