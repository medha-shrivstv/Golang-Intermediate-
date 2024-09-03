package main

import (
	"fmt"
)

type Point2D struct {
	X int
	Y int
}

func main() {
	p1 := Point2D{10, 50}
	fmt.Println("p1 = ", p1)
	p1 = Point2D{X: 10}
	fmt.Println("p1 = ", p1)
	p1 = Point2D{X: 310, Y: 50}
	fmt.Println("p1 = ", p1)
}
