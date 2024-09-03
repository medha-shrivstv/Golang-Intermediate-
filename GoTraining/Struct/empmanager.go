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

// EmpManager struct - contains a slice of Emp
type EmpManager struct {
	Employees []*Emp
}

// Add method - adds an employee to the slice
func (em *EmpManager) Add(emp *Emp) {
	em.Employees = append(em.Employees, emp)
}

// Print method - prints all employees in the slice
func (em EmpManager) Print() {
	for i, v := range em.Employees {
		fmt.Printf("Employee %d: ", i+1)
		v.Print()
	}
}

func main() {
	// Create some employees
	emp1 := &Emp{EmpNo: 1, EName: "John Doe", Salary: 50000}
	emp2 := &Emp{EmpNo: 2, EName: "Jane Smith", Salary: 60000}

	// Create EmpManager and add employees
	manager := EmpManager{}
	manager.Add(emp1)
	manager.Add(emp2)

	// Print all employees
	fmt.Println("Employees before salary increment:")
	manager.Print()

	// Increment salary for an employee
	emp1.SalaryIncr(10)
	emp2.SalaryIncr(20)

	// Print all employees after salary increment
	fmt.Println("\nEmployees after salary increment for the first employee:")
	manager.Print()
}
