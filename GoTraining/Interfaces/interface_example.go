package main

import "fmt"

type Display interface {
	Print()
}
type Dept struct {
	Deptno int
	Dname  string
}

func (d Dept) Print() {
	fmt.Println("in print of dept", d)
}

func main() {
	d1 := Dept{10, "Vaishali"}
	fmt.Println(d1)
	var display Display
	display = d1
	display.Print()
	fmt.Println(display)
}
