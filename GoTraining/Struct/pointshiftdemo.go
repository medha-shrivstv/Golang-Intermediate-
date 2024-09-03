package main

import "fmt"

type Point struct {
	X int
	Y int
}

func (p1 Point) print2() {
	fmt.Println("Print of Point ", p1)
}
func (p1 *Point) shift(dx, dy int) {
	p1.X += dx
	p1.Y += dy
	fmt.Println("p1 in shift = ", p1)
}
func main() {
	p1 := Point{10, 100}
	fmt.Println("p1 before shift ", p1)
	p1.shift(5, 5)
	fmt.Println("p1 after Shift ", p1)
	p1.print2()

}
