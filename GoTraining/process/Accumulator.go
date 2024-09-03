package main

import "fmt"

func main() {
	fp := process()
	no1 := fp(10)
	fmt.Println("No1 = ", no1)
	no1 = fp(5)
	fmt.Println("No1 = ", no1)
	no1 = fp(15)
	fmt.Println("No1 = ", no1)
}
func process() func(int) int {
	sum := 0
	return func(x int) int {
		fmt.Println("in func where sum =  ", sum)
		sum = sum + x
		return sum
	}
}
