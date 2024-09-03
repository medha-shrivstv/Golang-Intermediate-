package main

import (
	"fmt"

	"fands.com/calc"
)

func main() {
	ans := calc.Add(10, 20)
	fmt.Println("Sum = ", ans)
	ans = calc.Divide(50, 10)
	fmt.Println("Divide = ", ans)
}
