package main

import (
	"fmt"
)

// Emp struct definition
type Emp struct {
	EmpNo  int
	EName  string
	Salary float64
}

// Print method - prints the employee details
func (e Emp) Print() {
	fmt.Printf("EmpNo: %d, EName: %s, Salary: %.2f\n", e.EmpNo, e.EName, e.Salary)
}

// SalaryIncr method - increases the employee's salary by a given percentage
func (e *Emp) SalaryIncr(percent int) {
	increase := e.Salary * float64(percent) / 100
	e.Salary += increase
}

func main() {
	// Create an employee
	emp1 := Emp{EmpNo: 1, EName: "John Doe", Salary: 50000}

	// Print initial details
	fmt.Println("Before salary increment:")
	emp1.Print()

	// Increment salary by 10%
	emp1.SalaryIncr(10)

	// Print details after salary increment
	fmt.Println("\nAfter salary increment:")
	emp1.Print()
}
